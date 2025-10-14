package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
)

type UpdateHeartbeatCommand struct {
	AgentID agent.AgentID
}

type UpdateHeartbeatCommandHandler struct {
	repository agent.Repository
	eventBus   shared.EventBus
}

func NewUpdateHeartbeatCommandHandler(
	repository agent.Repository,
	eventBus shared.EventBus,
) *UpdateHeartbeatCommandHandler {
	return &UpdateHeartbeatCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *UpdateHeartbeatCommandHandler) Handle(ctx context.Context, cmd UpdateHeartbeatCommand) (*agent.Agent, error) {
	// Find agent
	a, err := h.repository.FindByID(ctx, cmd.AgentID)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, agent.ErrAgentNotFound
	}

	// Update heartbeat
	a.UpdateHeartbeat()

	// Save agent
	if err := h.repository.Save(ctx, a); err != nil {
		return nil, err
	}

	// Publish events
	for _, event := range a.Events {
		h.eventBus.Publish(event)
	}

	return a, nil
}
