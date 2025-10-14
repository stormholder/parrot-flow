package commands

// RegisterAgentRequest represents the request to register a new agent
type RegisterAgentRequest struct {
	Body struct {
		Name           string                  `json:"name" minLength:"1" maxLength:"255" doc:"Agent name"`
		Capabilities   CapabilitiesDTO         `json:"capabilities" doc:"Agent capabilities"`
		ConnectionInfo ConnectionInfoDTO       `json:"connection_info" doc:"Agent connection information"`
	}
}

// RegisterAgentResponse represents the response after registering an agent
type RegisterAgentResponse struct {
	Body struct {
		ID             string                  `json:"id" doc:"Agent ID"`
		Name           string                  `json:"name" doc:"Agent name"`
		Status         string                  `json:"status" doc:"Agent status"`
		Capabilities   CapabilitiesDTO         `json:"capabilities" doc:"Agent capabilities"`
		ConnectionInfo ConnectionInfoDTO       `json:"connection_info" doc:"Agent connection information"`
		RegisteredAt   string                  `json:"registered_at" doc:"Registration timestamp"`
	}
}

// UpdateHeartbeatRequest represents the request to update agent heartbeat
type UpdateHeartbeatRequest struct {
	ID string `path:"id" doc:"Agent ID"`
}

// UpdateHeartbeatResponse represents the response after updating heartbeat
type UpdateHeartbeatResponse struct {
	Body struct {
		ID              string `json:"id" doc:"Agent ID"`
		Name            string `json:"name" doc:"Agent name"`
		Status          string `json:"status" doc:"Agent status"`
		LastHeartbeatAt string `json:"last_heartbeat_at" doc:"Last heartbeat timestamp"`
	}
}

// AssignRunRequest represents the request to assign a run to an agent
type AssignRunRequest struct {
	ID string `path:"id" doc:"Agent ID"`
}

// AssignRunResponse represents the response after assigning a run
type AssignRunResponse struct {
	Body struct {
		ID              string `json:"id" doc:"Agent ID"`
		Name            string `json:"name" doc:"Agent name"`
		Status          string `json:"status" doc:"Agent status"`
		CurrentRunCount int    `json:"current_run_count" doc:"Current number of runs"`
	}
}

// ReleaseRunRequest represents the request to release a run from an agent
type ReleaseRunRequest struct {
	ID string `path:"id" doc:"Agent ID"`
}

// ReleaseRunResponse represents the response after releasing a run
type ReleaseRunResponse struct {
	Body struct {
		ID              string `json:"id" doc:"Agent ID"`
		Name            string `json:"name" doc:"Agent name"`
		Status          string `json:"status" doc:"Agent status"`
		CurrentRunCount int    `json:"current_run_count" doc:"Current number of runs"`
	}
}

// UpdateAgentRequest represents the request to update an agent
type UpdateAgentRequest struct {
	ID   string `path:"id" doc:"Agent ID"`
	Body struct {
		Name         *string          `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Agent name"`
		Capabilities *CapabilitiesDTO `json:"capabilities,omitempty" doc:"Agent capabilities"`
		TagsToAdd    []string         `json:"tags_to_add,omitempty" doc:"Tag IDs to add"`
		TagsToRemove []string         `json:"tags_to_remove,omitempty" doc:"Tag IDs to remove"`
	}
}

// UpdateAgentResponse represents the response after updating an agent
type UpdateAgentResponse struct {
	Body struct {
		ID           string          `json:"id" doc:"Agent ID"`
		Name         string          `json:"name" doc:"Agent name"`
		Status       string          `json:"status" doc:"Agent status"`
		Capabilities CapabilitiesDTO `json:"capabilities" doc:"Agent capabilities"`
		Tags         []string        `json:"tags" doc:"Tag IDs"`
		UpdatedAt    string          `json:"updated_at" doc:"Update timestamp"`
	}
}

// DeregisterAgentRequest represents the request to deregister an agent
type DeregisterAgentRequest struct {
	ID string `path:"id" doc:"Agent ID"`
}

// DeregisterAgentResponse represents the response after deregistering an agent
type DeregisterAgentResponse struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	}
}

// CapabilitiesDTO represents agent capabilities in API format
type CapabilitiesDTO struct {
	Browsers []BrowserCapabilityDTO `json:"browsers" doc:"Supported browsers"`
	OS       OSInfoDTO              `json:"os" doc:"Operating system information"`
	Proxy    ProxyCapabilityDTO     `json:"proxy" doc:"Proxy support"`
	Resource ResourceLimitsDTO      `json:"resource" doc:"Resource limits"`
	Features []string               `json:"features,omitempty" doc:"Additional features"`
}

// BrowserCapabilityDTO represents a browser capability in API format
type BrowserCapabilityDTO struct {
	Type                string   `json:"type" enum:"chromium,firefox,webkit" doc:"Browser type"`
	Version             string   `json:"version" doc:"Browser version"`
	Headless            bool     `json:"headless" doc:"Headless support"`
	SupportedExtensions []string `json:"supported_extensions,omitempty" doc:"Supported extensions"`
}

// OSInfoDTO represents operating system information in API format
type OSInfoDTO struct {
	Platform     string `json:"platform" enum:"linux,darwin,windows" doc:"Operating system platform"`
	Architecture string `json:"architecture" enum:"amd64,arm64,386,arm" doc:"CPU architecture"`
	Version      string `json:"version" doc:"OS version"`
}

// ProxyCapabilityDTO represents proxy capability in API format
type ProxyCapabilityDTO struct {
	SupportsProxy      bool     `json:"supports_proxy" doc:"Whether proxy is supported"`
	SupportedProtocols []string `json:"supported_protocols,omitempty" doc:"Supported proxy protocols"`
}

// ResourceLimitsDTO represents resource limits in API format
type ResourceLimitsDTO struct {
	MaxConcurrentRuns int `json:"max_concurrent_runs" minimum:"1" doc:"Maximum concurrent runs"`
	MaxMemoryMB       int `json:"max_memory_mb" minimum:"1" doc:"Maximum memory in MB"`
	MaxCPUCores       int `json:"max_cpu_cores" minimum:"1" doc:"Maximum CPU cores"`
}

// ConnectionInfoDTO represents connection information in API format
type ConnectionInfoDTO struct {
	IPAddress string `json:"ip_address,omitempty" doc:"Agent IP address"`
	Hostname  string `json:"hostname,omitempty" doc:"Agent hostname"`
	QueueName string `json:"queue_name" doc:"RabbitMQ queue name"`
}
