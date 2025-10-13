package agent

import (
	"parrotflow/internal/domain/shared"
	"time"
)

const (
	EventAgentRegistered           = "agent.registered"
	EventAgentDeregistered         = "agent.deregistered"
	EventAgentStatusChanged        = "agent.status.changed"
	EventAgentDisconnected         = "agent.disconnected"
	EventAgentCapabilitiesUpdated  = "agent.capabilities.updated"
)

type AgentRegistered struct {
	shared.BaseEvent
	AgentID      string
	Name         string
	RegisteredAt time.Time
}

type AgentDeregistered struct {
	shared.BaseEvent
	AgentID        string
	DeregisteredAt time.Time
}

type AgentStatusChanged struct {
	shared.BaseEvent
	AgentID   string
	OldStatus string
	NewStatus string
}

type AgentDisconnected struct {
	shared.BaseEvent
	AgentID        string
	DisconnectedAt time.Time
}

type AgentCapabilitiesUpdated struct {
	shared.BaseEvent
	AgentID string
}
