package agent

import (
	"errors"
	"fmt"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
	"time"
)

type AgentID struct {
	shared.ID
}

func NewAgentID(value string) (AgentID, error) {
	id, err := shared.NewID(value)
	if err != nil {
		return AgentID{}, err
	}
	return AgentID{ID: id}, nil
}

// AgentStatus represents the current status of an agent
type AgentStatus struct {
	value string
}

func NewAgentStatus(value string) (AgentStatus, error) {
	switch value {
	case "online", "offline", "busy", "idle", "disconnected":
		return AgentStatus{value: value}, nil
	default:
		return AgentStatus{}, fmt.Errorf("invalid agent status: %s", value)
	}
}

func (as AgentStatus) String() string {
	return as.value
}

// Common agent statuses
var (
	AgentStatusOnline       = AgentStatus{value: "online"}
	AgentStatusOffline      = AgentStatus{value: "offline"}
	AgentStatusBusy         = AgentStatus{value: "busy"}
	AgentStatusIdle         = AgentStatus{value: "idle"}
	AgentStatusDisconnected = AgentStatus{value: "disconnected"}
)

// Agent represents a browser automation agent
type Agent struct {
	Id                AgentID
	Name              string
	Status            AgentStatus
	Capabilities      Capabilities
	Tags              []tag.TagID
	CurrentRunCount   int
	LastHeartbeatAt   *shared.Timestamp
	RegisteredAt      shared.Timestamp
	UpdatedAt         shared.Timestamp
	ConnectionInfo    ConnectionInfo
	Metadata          map[string]interface{} // Extensible metadata
	Events            []shared.DomainEvent
}

// NewAgent creates a new agent (used for auto-registration)
func NewAgent(
	id AgentID,
	name string,
	capabilities Capabilities,
	connectionInfo ConnectionInfo,
) (*Agent, error) {
	if name == "" {
		return nil, errors.New("agent name cannot be empty")
	}

	now := shared.NewTimestamp(time.Now())
	agent := &Agent{
		Id:               id,
		Name:             name,
		Status:           AgentStatusOnline,
		Capabilities:     capabilities,
		Tags:             make([]tag.TagID, 0),
		CurrentRunCount:  0,
		LastHeartbeatAt:  &now,
		RegisteredAt:     now,
		UpdatedAt:        now,
		ConnectionInfo:   connectionInfo,
		Metadata:         make(map[string]interface{}),
		Events:           make([]shared.DomainEvent, 0),
	}

	agent.addEvent(AgentRegistered{
		BaseEvent:    shared.NewBaseEvent(EventAgentRegistered, id.String()),
		AgentID:      id.String(),
		Name:         name,
		RegisteredAt: now.Time(),
	})

	return agent, nil
}

// UpdateHeartbeat updates the agent's last heartbeat time
func (a *Agent) UpdateHeartbeat() {
	now := shared.NewTimestamp(time.Now())
	a.LastHeartbeatAt = &now
	a.UpdatedAt = shared.NewTimestamp(time.Now())

	// Auto-recover from disconnected status on heartbeat
	if a.Status.String() == AgentStatusDisconnected.String() {
		a.UpdateStatus(AgentStatusIdle)
	}
}

// UpdateStatus updates the agent status
func (a *Agent) UpdateStatus(status AgentStatus) {
	oldStatus := a.Status
	a.Status = status
	a.UpdatedAt = shared.NewTimestamp(time.Now())

	if oldStatus.String() != status.String() {
		a.addEvent(AgentStatusChanged{
			BaseEvent: shared.NewBaseEvent(EventAgentStatusChanged, a.Id.String()),
			AgentID:   a.Id.String(),
			OldStatus: oldStatus.String(),
			NewStatus: status.String(),
		})
	}
}

// MarkDisconnected marks the agent as disconnected (called when heartbeat timeout occurs)
func (a *Agent) MarkDisconnected() {
	a.UpdateStatus(AgentStatusDisconnected)
	a.addEvent(AgentDisconnected{
		BaseEvent:      shared.NewBaseEvent(EventAgentDisconnected, a.Id.String()),
		AgentID:        a.Id.String(),
		DisconnectedAt: time.Now(),
	})
}

// CanAcceptRun checks if the agent can accept a new run
func (a *Agent) CanAcceptRun() bool {
	if a.Status.String() == AgentStatusDisconnected.String() ||
		a.Status.String() == AgentStatusOffline.String() {
		return false
	}
	return a.CurrentRunCount < a.Capabilities.ResourceLimits.MaxConcurrentRuns
}

// AssignRun assigns a run to the agent
func (a *Agent) AssignRun() error {
	if !a.CanAcceptRun() {
		return fmt.Errorf("agent cannot accept run: status=%s, current=%d, max=%d",
			a.Status.String(),
			a.CurrentRunCount,
			a.Capabilities.ResourceLimits.MaxConcurrentRuns)
	}

	a.CurrentRunCount++
	a.UpdatedAt = shared.NewTimestamp(time.Now())

	// Update status to busy if at capacity
	if a.CurrentRunCount >= a.Capabilities.ResourceLimits.MaxConcurrentRuns {
		a.UpdateStatus(AgentStatusBusy)
	} else if a.Status.String() == AgentStatusIdle.String() {
		a.UpdateStatus(AgentStatusOnline)
	}

	return nil
}

// ReleaseRun releases a run from the agent
func (a *Agent) ReleaseRun() error {
	if a.CurrentRunCount <= 0 {
		return errors.New("no runs to release")
	}

	a.CurrentRunCount--
	a.UpdatedAt = shared.NewTimestamp(time.Now())

	// Update status if no longer busy
	if a.CurrentRunCount < a.Capabilities.ResourceLimits.MaxConcurrentRuns {
		if a.Status.String() == AgentStatusBusy.String() {
			a.UpdateStatus(AgentStatusIdle)
		}
	}

	return nil
}

// AddTag adds a tag to the agent
func (a *Agent) AddTag(tagID tag.TagID) error {
	// Check if tag already exists
	for _, existingTag := range a.Tags {
		if existingTag.String() == tagID.String() {
			return errors.New("tag already exists on agent")
		}
	}
	a.Tags = append(a.Tags, tagID)
	a.UpdatedAt = shared.NewTimestamp(time.Now())
	return nil
}

// RemoveTag removes a tag from the agent
func (a *Agent) RemoveTag(tagID tag.TagID) {
	for i, existingTag := range a.Tags {
		if existingTag.String() == tagID.String() {
			a.Tags = append(a.Tags[:i], a.Tags[i+1:]...)
			a.UpdatedAt = shared.NewTimestamp(time.Now())
			return
		}
	}
}

// HasTag checks if the agent has a specific tag
func (a *Agent) HasTag(tagID tag.TagID) bool {
	for _, existingTag := range a.Tags {
		if existingTag.String() == tagID.String() {
			return true
		}
	}
	return false
}

// HasAllTags checks if the agent has all specified tags
func (a *Agent) HasAllTags(tagIDs []tag.TagID) bool {
	for _, requiredTag := range tagIDs {
		if !a.HasTag(requiredTag) {
			return false
		}
	}
	return true
}

// UpdateCapabilities updates the agent's capabilities
func (a *Agent) UpdateCapabilities(capabilities Capabilities) {
	a.Capabilities = capabilities
	a.UpdatedAt = shared.NewTimestamp(time.Now())

	a.addEvent(AgentCapabilitiesUpdated{
		BaseEvent: shared.NewBaseEvent(EventAgentCapabilitiesUpdated, a.Id.String()),
		AgentID:   a.Id.String(),
	})
}

// UpdateName updates the agent's name
func (a *Agent) UpdateName(name string) error {
	if name == "" {
		return errors.New("agent name cannot be empty")
	}
	a.Name = name
	a.UpdatedAt = shared.NewTimestamp(time.Now())
	return nil
}

// SetMetadata sets a metadata value
func (a *Agent) SetMetadata(key string, value interface{}) {
	a.Metadata[key] = value
	a.UpdatedAt = shared.NewTimestamp(time.Now())
}

// GetMetadata gets a metadata value
func (a *Agent) GetMetadata(key string) (interface{}, bool) {
	value, exists := a.Metadata[key]
	return value, exists
}

// IsHealthy checks if the agent is healthy based on last heartbeat
func (a *Agent) IsHealthy(heartbeatTimeout time.Duration) bool {
	if a.LastHeartbeatAt == nil {
		return false
	}
	timeSinceHeartbeat := time.Since(a.LastHeartbeatAt.Time())
	return timeSinceHeartbeat < heartbeatTimeout
}

// Deregister marks the agent for deregistration
func (a *Agent) Deregister() {
	a.UpdateStatus(AgentStatusOffline)
	a.addEvent(AgentDeregistered{
		BaseEvent:      shared.NewBaseEvent(EventAgentDeregistered, a.Id.String()),
		AgentID:        a.Id.String(),
		DeregisteredAt: time.Now(),
	})
}

func (a *Agent) addEvent(event shared.DomainEvent) {
	a.Events = append(a.Events, event)
}

func (a *Agent) ClearEvents() {
	a.Events = make([]shared.DomainEvent, 0)
}
