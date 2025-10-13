package ports

import (
	"encoding/json"
	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/models"
)

func AgentParseID(id string) uint64 {
	return parseID(id)
}

// CapabilitiesDTO represents capabilities in JSON format
type CapabilitiesDTO struct {
	Browsers []BrowserCapabilityDTO `json:"browsers"`
	OS       OSInfoDTO              `json:"os"`
	Proxy    ProxyCapabilityDTO     `json:"proxy"`
	Resource ResourceLimitsDTO      `json:"resource_limits"`
	Features []string               `json:"features"`
}

type BrowserCapabilityDTO struct {
	Type                string   `json:"type"`
	Version             string   `json:"version"`
	Headless            bool     `json:"headless"`
	SupportedExtensions []string `json:"supported_extensions"`
}

type OSInfoDTO struct {
	Platform     string `json:"platform"`
	Architecture string `json:"architecture"`
	Version      string `json:"version"`
}

type ProxyCapabilityDTO struct {
	SupportsProxy      bool     `json:"supports_proxy"`
	SupportedProtocols []string `json:"supported_protocols"`
}

type ResourceLimitsDTO struct {
	MaxConcurrentRuns int `json:"max_concurrent_runs"`
	MaxMemoryMB       int `json:"max_memory_mb"`
	MaxCPUCores       int `json:"max_cpu_cores"`
}

// ConnectionInfoDTO represents connection info in JSON format
type ConnectionInfoDTO struct {
	IPAddress string `json:"ip_address"`
	Hostname  string `json:"hostname"`
	QueueName string `json:"queue_name"`
}

func AgentDomainEntityToPersistence(a *agent.Agent) (*models.Agent, error) {
	model := &models.Agent{
		Model: models.Model{
			ID:        parseID(a.Id.String()),
			CreatedAt: a.RegisteredAt.Time(),
			UpdatedAt: a.UpdatedAt.Time(),
		},
		Name:            a.Name,
		Status:          a.Status.String(),
		CurrentRunCount: a.CurrentRunCount,
		RegisteredAt:    a.RegisteredAt.Time(),
	}

	// Marshal capabilities
	capabilitiesJSON, err := marshalCapabilities(a.Capabilities)
	if err != nil {
		return nil, err
	}
	model.Capabilities = capabilitiesJSON

	// Marshal connection info
	connInfoJSON, err := marshalConnectionInfo(a.ConnectionInfo)
	if err != nil {
		return nil, err
	}
	model.ConnectionInfo = connInfoJSON

	// Marshal metadata
	if len(a.Metadata) > 0 {
		metadataJSON, err := json.Marshal(a.Metadata)
		if err != nil {
			return nil, err
		}
		model.Metadata = string(metadataJSON)
	}

	// Handle heartbeat timestamp
	if a.LastHeartbeatAt != nil {
		t := a.LastHeartbeatAt.Time()
		model.LastHeartbeatAt = &t
	}

	return model, nil
}

func AgentPersistenceToDomainEntity(model *models.Agent) (*agent.Agent, error) {
	agentID, err := agent.NewAgentID(formatID(model.ID))
	if err != nil {
		return nil, err
	}

	// Unmarshal capabilities
	capabilities, err := unmarshalCapabilities(model.Capabilities)
	if err != nil {
		return nil, err
	}

	// Unmarshal connection info
	connInfo, err := unmarshalConnectionInfo(model.ConnectionInfo)
	if err != nil {
		return nil, err
	}

	// Create agent
	a, err := agent.NewAgent(agentID, model.Name, capabilities, connInfo)
	if err != nil {
		return nil, err
	}

	// Set status
	status, err := agent.NewAgentStatus(model.Status)
	if err != nil {
		return nil, err
	}
	a.UpdateStatus(status)

	// Set run count
	a.CurrentRunCount = model.CurrentRunCount

	// Set heartbeat
	if model.LastHeartbeatAt != nil {
		ts := shared.NewTimestamp(*model.LastHeartbeatAt)
		a.LastHeartbeatAt = &ts
	}

	// Unmarshal metadata
	if model.Metadata != "" {
		var metadata map[string]interface{}
		if err := json.Unmarshal([]byte(model.Metadata), &metadata); err != nil {
			return nil, err
		}
		a.Metadata = metadata
	}

	// Convert tags
	for _, tagModel := range model.Tags {
		tagID, err := tag.NewTagID(formatID(tagModel.ID))
		if err != nil {
			return nil, err
		}
		a.AddTag(tagID)
	}

	return a, nil
}

func marshalCapabilities(cap agent.Capabilities) (string, error) {
	dto := CapabilitiesDTO{
		Browsers: make([]BrowserCapabilityDTO, len(cap.Browsers)),
		OS: OSInfoDTO{
			Platform:     cap.OS.Platform.String(),
			Architecture: cap.OS.Architecture.String(),
			Version:      cap.OS.Version,
		},
		Proxy: ProxyCapabilityDTO{
			SupportsProxy:      cap.Proxy.SupportsProxy,
			SupportedProtocols: cap.Proxy.SupportedProtocols,
		},
		Resource: ResourceLimitsDTO{
			MaxConcurrentRuns: cap.ResourceLimits.MaxConcurrentRuns,
			MaxMemoryMB:       cap.ResourceLimits.MaxMemoryMB,
			MaxCPUCores:       cap.ResourceLimits.MaxCPUCores,
		},
		Features: cap.Features,
	}

	for i, browser := range cap.Browsers {
		dto.Browsers[i] = BrowserCapabilityDTO{
			Type:                browser.Type.String(),
			Version:             browser.Version,
			Headless:            browser.Headless,
			SupportedExtensions: browser.SupportedExtensions,
		}
	}

	data, err := json.Marshal(dto)
	return string(data), err
}

func unmarshalCapabilities(data string) (agent.Capabilities, error) {
	var dto CapabilitiesDTO
	if err := json.Unmarshal([]byte(data), &dto); err != nil {
		return agent.Capabilities{}, err
	}

	// Convert browsers
	browsers := make([]agent.BrowserCapability, len(dto.Browsers))
	for i, b := range dto.Browsers {
		browserType, err := agent.NewBrowserType(b.Type)
		if err != nil {
			return agent.Capabilities{}, err
		}
		browserCap, err := agent.NewBrowserCapability(browserType, b.Version, b.Headless)
		if err != nil {
			return agent.Capabilities{}, err
		}
		browserCap.SupportedExtensions = b.SupportedExtensions
		browsers[i] = browserCap
	}

	// Convert OS
	platform, err := agent.NewPlatform(dto.OS.Platform)
	if err != nil {
		return agent.Capabilities{}, err
	}
	arch, err := agent.NewArchitecture(dto.OS.Architecture)
	if err != nil {
		return agent.Capabilities{}, err
	}
	osInfo, err := agent.NewOSInfo(platform, arch, dto.OS.Version)
	if err != nil {
		return agent.Capabilities{}, err
	}

	// Convert proxy capability
	proxyCapability := agent.NewProxyCapability(dto.Proxy.SupportsProxy, dto.Proxy.SupportedProtocols)

	// Convert resource limits
	resourceLimits, err := agent.NewResourceLimits(
		dto.Resource.MaxConcurrentRuns,
		dto.Resource.MaxMemoryMB,
		dto.Resource.MaxCPUCores,
	)
	if err != nil {
		return agent.Capabilities{}, err
	}

	// Create capabilities
	return agent.NewCapabilities(browsers, osInfo, proxyCapability, resourceLimits, dto.Features)
}

func marshalConnectionInfo(connInfo agent.ConnectionInfo) (string, error) {
	dto := ConnectionInfoDTO{
		IPAddress: connInfo.IPAddress,
		Hostname:  connInfo.Hostname,
		QueueName: connInfo.QueueName,
	}
	data, err := json.Marshal(dto)
	return string(data), err
}

func unmarshalConnectionInfo(data string) (agent.ConnectionInfo, error) {
	var dto ConnectionInfoDTO
	if err := json.Unmarshal([]byte(data), &dto); err != nil {
		return agent.ConnectionInfo{}, err
	}
	return agent.NewConnectionInfo(dto.IPAddress, dto.Hostname, dto.QueueName)
}
