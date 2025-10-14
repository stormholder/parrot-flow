package mappers

import (
	"parrotflow/internal/domain/agent"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

// ToRegisterAgentResponse converts domain agent to register response DTO
func ToRegisterAgentResponse(a *agent.Agent) *commands.RegisterAgentResponse {
	response := &commands.RegisterAgentResponse{}
	response.Body.ID = a.Id.String()
	response.Body.Name = a.Name
	response.Body.Status = a.Status.String()
	response.Body.Capabilities = ToCapabilitiesDTO(a.Capabilities)
	response.Body.ConnectionInfo = ToConnectionInfoDTO(a.ConnectionInfo)
	response.Body.RegisteredAt = a.RegisteredAt.Time().Format("2006-01-02T15:04:05Z07:00")
	return response
}

// ToUpdateHeartbeatResponse converts domain agent to heartbeat response DTO
func ToUpdateHeartbeatResponse(a *agent.Agent) *commands.UpdateHeartbeatResponse {
	response := &commands.UpdateHeartbeatResponse{}
	response.Body.ID = a.Id.String()
	response.Body.Name = a.Name
	response.Body.Status = a.Status.String()
	if a.LastHeartbeatAt != nil {
		response.Body.LastHeartbeatAt = a.LastHeartbeatAt.Time().Format("2006-01-02T15:04:05Z07:00")
	}
	return response
}

// ToAssignRunResponse converts domain agent to assign run response DTO
func ToAssignRunResponse(a *agent.Agent) *commands.AssignRunResponse {
	response := &commands.AssignRunResponse{}
	response.Body.ID = a.Id.String()
	response.Body.Name = a.Name
	response.Body.Status = a.Status.String()
	response.Body.CurrentRunCount = a.CurrentRunCount
	return response
}

// ToReleaseRunResponse converts domain agent to release run response DTO
func ToReleaseRunResponse(a *agent.Agent) *commands.ReleaseRunResponse {
	response := &commands.ReleaseRunResponse{}
	response.Body.ID = a.Id.String()
	response.Body.Name = a.Name
	response.Body.Status = a.Status.String()
	response.Body.CurrentRunCount = a.CurrentRunCount
	return response
}

// ToUpdateAgentResponse converts domain agent to update response DTO
func ToUpdateAgentResponse(a *agent.Agent) *commands.UpdateAgentResponse {
	response := &commands.UpdateAgentResponse{}
	response.Body.ID = a.Id.String()
	response.Body.Name = a.Name
	response.Body.Status = a.Status.String()
	response.Body.Capabilities = ToCapabilitiesDTO(a.Capabilities)
	response.Body.Tags = make([]string, len(a.Tags))
	for i, tagID := range a.Tags {
		response.Body.Tags[i] = tagID.String()
	}
	response.Body.UpdatedAt = a.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	return response
}

// ToDeregisterAgentResponse creates deregister response DTO
func ToDeregisterAgentResponse() *commands.DeregisterAgentResponse {
	response := &commands.DeregisterAgentResponse{}
	response.Body.Message = "Agent deregistered successfully"
	return response
}

// ToGetAgentResponse converts domain agent to get response DTO
func ToGetAgentResponse(a *agent.Agent) *queries.GetAgentResponse {
	response := &queries.GetAgentResponse{}
	response.Body = ToAgentDTO(a)
	return response
}

// ToListAgentsResponse converts domain agents to list response DTO
func ToListAgentsResponse(agents []*agent.Agent) *queries.ListAgentsResponse {
	response := &queries.ListAgentsResponse{}
	response.Body.Agents = make([]queries.AgentDTO, len(agents))
	for i, a := range agents {
		response.Body.Agents[i] = ToAgentDTO(a)
	}
	response.Body.Count = len(agents)
	return response
}

// ToGetAvailableAgentsResponse converts domain agents to available agents response DTO
func ToGetAvailableAgentsResponse(agents []*agent.Agent) *queries.GetAvailableAgentsResponse {
	response := &queries.GetAvailableAgentsResponse{}
	response.Body.Agents = make([]queries.AgentDTO, len(agents))
	for i, a := range agents {
		response.Body.Agents[i] = ToAgentDTO(a)
	}
	response.Body.Count = len(agents)
	return response
}

// ToGetStaleAgentsResponse converts domain agents to stale agents response DTO
func ToGetStaleAgentsResponse(agents []*agent.Agent) *queries.GetStaleAgentsResponse {
	response := &queries.GetStaleAgentsResponse{}
	response.Body.Agents = make([]queries.AgentDTO, len(agents))
	for i, a := range agents {
		response.Body.Agents[i] = ToAgentDTO(a)
	}
	response.Body.Count = len(agents)
	return response
}

// ToAgentDTO converts domain agent to DTO
func ToAgentDTO(a *agent.Agent) queries.AgentDTO {
	dto := queries.AgentDTO{
		ID:              a.Id.String(),
		Name:            a.Name,
		Status:          a.Status.String(),
		Capabilities:    ToCapabilitiesDTO(a.Capabilities),
		Tags:            make([]string, len(a.Tags)),
		CurrentRunCount: a.CurrentRunCount,
		RegisteredAt:    a.RegisteredAt.Time().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:       a.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00"),
		ConnectionInfo:  ToConnectionInfoDTO(a.ConnectionInfo),
	}

	// Convert tags
	for i, tagID := range a.Tags {
		dto.Tags[i] = tagID.String()
	}

	// Convert last heartbeat if present
	if a.LastHeartbeatAt != nil {
		heartbeat := a.LastHeartbeatAt.Time().Format("2006-01-02T15:04:05Z07:00")
		dto.LastHeartbeatAt = &heartbeat
	}

	return dto
}

// ToCapabilitiesDTO converts domain capabilities to DTO
func ToCapabilitiesDTO(c agent.Capabilities) commands.CapabilitiesDTO {
	dto := commands.CapabilitiesDTO{
		Browsers: make([]commands.BrowserCapabilityDTO, len(c.Browsers)),
		OS:       ToOSInfoDTO(c.OS),
		Proxy:    ToProxyCapabilityDTO(c.Proxy),
		Resource: ToResourceLimitsDTO(c.ResourceLimits),
		Features: c.Features,
	}

	for i, browser := range c.Browsers {
		dto.Browsers[i] = ToBrowserCapabilityDTO(browser)
	}

	return dto
}

// ToBrowserCapabilityDTO converts domain browser capability to DTO
func ToBrowserCapabilityDTO(b agent.BrowserCapability) commands.BrowserCapabilityDTO {
	return commands.BrowserCapabilityDTO{
		Type:                b.Type.String(),
		Version:             b.Version,
		Headless:            b.Headless,
		SupportedExtensions: b.SupportedExtensions,
	}
}

// ToOSInfoDTO converts domain OS info to DTO
func ToOSInfoDTO(o agent.OSInfo) commands.OSInfoDTO {
	return commands.OSInfoDTO{
		Platform:     o.Platform.String(),
		Architecture: o.Architecture.String(),
		Version:      o.Version,
	}
}

// ToProxyCapabilityDTO converts domain proxy capability to DTO
func ToProxyCapabilityDTO(p agent.ProxyCapability) commands.ProxyCapabilityDTO {
	return commands.ProxyCapabilityDTO{
		SupportsProxy:      p.SupportsProxy,
		SupportedProtocols: p.SupportedProtocols,
	}
}

// ToResourceLimitsDTO converts domain resource limits to DTO
func ToResourceLimitsDTO(r agent.ResourceLimits) commands.ResourceLimitsDTO {
	return commands.ResourceLimitsDTO{
		MaxConcurrentRuns: r.MaxConcurrentRuns,
		MaxMemoryMB:       r.MaxMemoryMB,
		MaxCPUCores:       r.MaxCPUCores,
	}
}

// ToConnectionInfoDTO converts domain connection info to DTO
func ToConnectionInfoDTO(c agent.ConnectionInfo) commands.ConnectionInfoDTO {
	return commands.ConnectionInfoDTO{
		IPAddress: c.IPAddress,
		Hostname:  c.Hostname,
		QueueName: c.QueueName,
	}
}

// FromCapabilitiesDTO converts DTO to domain capabilities
func FromCapabilitiesDTO(dto commands.CapabilitiesDTO) (agent.Capabilities, error) {
	browsers := make([]agent.BrowserCapability, len(dto.Browsers))
	for i, b := range dto.Browsers {
		browser, err := FromBrowserCapabilityDTO(b)
		if err != nil {
			return agent.Capabilities{}, err
		}
		browsers[i] = browser
	}

	osInfo, err := FromOSInfoDTO(dto.OS)
	if err != nil {
		return agent.Capabilities{}, err
	}

	proxyCapability := FromProxyCapabilityDTO(dto.Proxy)

	resourceLimits, err := FromResourceLimitsDTO(dto.Resource)
	if err != nil {
		return agent.Capabilities{}, err
	}

	return agent.NewCapabilities(browsers, osInfo, proxyCapability, resourceLimits, dto.Features)
}

// FromBrowserCapabilityDTO converts DTO to domain browser capability
func FromBrowserCapabilityDTO(dto commands.BrowserCapabilityDTO) (agent.BrowserCapability, error) {
	browserType, err := agent.NewBrowserType(dto.Type)
	if err != nil {
		return agent.BrowserCapability{}, err
	}

	capability, err := agent.NewBrowserCapability(browserType, dto.Version, dto.Headless)
	if err != nil {
		return agent.BrowserCapability{}, err
	}

	if dto.SupportedExtensions != nil {
		capability.SupportedExtensions = dto.SupportedExtensions
	}

	return capability, nil
}

// FromOSInfoDTO converts DTO to domain OS info
func FromOSInfoDTO(dto commands.OSInfoDTO) (agent.OSInfo, error) {
	platform, err := agent.NewPlatform(dto.Platform)
	if err != nil {
		return agent.OSInfo{}, err
	}

	architecture, err := agent.NewArchitecture(dto.Architecture)
	if err != nil {
		return agent.OSInfo{}, err
	}

	return agent.NewOSInfo(platform, architecture, dto.Version)
}

// FromProxyCapabilityDTO converts DTO to domain proxy capability
func FromProxyCapabilityDTO(dto commands.ProxyCapabilityDTO) agent.ProxyCapability {
	return agent.NewProxyCapability(dto.SupportsProxy, dto.SupportedProtocols)
}

// FromResourceLimitsDTO converts DTO to domain resource limits
func FromResourceLimitsDTO(dto commands.ResourceLimitsDTO) (agent.ResourceLimits, error) {
	return agent.NewResourceLimits(dto.MaxConcurrentRuns, dto.MaxMemoryMB, dto.MaxCPUCores)
}

// FromConnectionInfoDTO converts DTO to domain connection info
func FromConnectionInfoDTO(dto commands.ConnectionInfoDTO) (agent.ConnectionInfo, error) {
	return agent.NewConnectionInfo(dto.IPAddress, dto.Hostname, dto.QueueName)
}
