package queries

import "parrotflow/internal/interfaces/http/dto/commands"

// GetAgentRequest represents the request to get an agent by ID
type GetAgentRequest struct {
	ID string `path:"id" doc:"Agent ID"`
}

// GetAgentResponse represents the response containing a single agent
type GetAgentResponse struct {
	Body AgentDTO `json:"agent" doc:"Agent details"`
}

// ListAgentsRequest represents the request to list agents with optional filters
type ListAgentsRequest struct {
	Status       string `query:"status" required:"false" enum:"online,offline,busy,idle,disconnected" doc:"Filter by status"`
	BrowserType  string `query:"browser_type" required:"false" enum:"chromium,firefox,webkit" doc:"Filter by browser type"`
	Platform     string `query:"platform" required:"false" enum:"linux,darwin,windows" doc:"Filter by platform"`
	OnlyHealthy  bool   `query:"only_healthy" required:"false" doc:"Filter only healthy agents"`
	TagIDs       string `query:"tag_ids" required:"false" doc:"Comma-separated tag IDs"`
}

// ListAgentsResponse represents the response containing a list of agents
type ListAgentsResponse struct {
	Body struct {
		Agents []AgentDTO `json:"agents" doc:"List of agents"`
		Count  int        `json:"count" doc:"Total number of agents"`
	}
}

// GetAvailableAgentsRequest represents the request to get available agents
type GetAvailableAgentsRequest struct {
	BrowserType string `query:"browser_type" required:"false" enum:"chromium,firefox,webkit" doc:"Filter by browser type"`
	Platform    string `query:"platform" required:"false" enum:"linux,darwin,windows" doc:"Filter by platform"`
	TagIDs      string `query:"tag_ids" required:"false" doc:"Comma-separated tag IDs"`
}

// GetAvailableAgentsResponse represents the response containing available agents
type GetAvailableAgentsResponse struct {
	Body struct {
		Agents []AgentDTO `json:"agents" doc:"List of available agents"`
		Count  int        `json:"count" doc:"Total number of available agents"`
	}
}

// GetStaleAgentsRequest represents the request to get stale agents
type GetStaleAgentsRequest struct {
	HeartbeatTimeoutMinutes int `query:"heartbeat_timeout_minutes" required:"false" minimum:"1" default:"5" doc:"Heartbeat timeout in minutes"`
}

// GetStaleAgentsResponse represents the response containing stale agents
type GetStaleAgentsResponse struct {
	Body struct {
		Agents []AgentDTO `json:"agents" doc:"List of stale agents"`
		Count  int        `json:"count" doc:"Total number of stale agents"`
	}
}

// AgentDTO represents an agent in API format
type AgentDTO struct {
	ID              string                       `json:"id" doc:"Agent ID"`
	Name            string                       `json:"name" doc:"Agent name"`
	Status          string                       `json:"status" doc:"Agent status"`
	Capabilities    commands.CapabilitiesDTO     `json:"capabilities" doc:"Agent capabilities"`
	Tags            []string                     `json:"tags" doc:"Tag IDs"`
	CurrentRunCount int                          `json:"current_run_count" doc:"Current number of runs"`
	LastHeartbeatAt *string                      `json:"last_heartbeat_at,omitempty" doc:"Last heartbeat timestamp"`
	RegisteredAt    string                       `json:"registered_at" doc:"Registration timestamp"`
	UpdatedAt       string                       `json:"updated_at" doc:"Update timestamp"`
	ConnectionInfo  commands.ConnectionInfoDTO   `json:"connection_info" doc:"Connection information"`
}
