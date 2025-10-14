package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
)

type DeregisterAgentCommand struct {
	AgentID agent.AgentID
}

type DeregisterAgentCommandHandler struct {
	repository agent.Repository
	eventBus   shared.EventBus
}

func NewDeregisterAgentCommandHandler(
	repository agent.Repository,
	eventBus shared.EventBus,
) *DeregisterAgentCommandHandler {
	return &DeregisterAgentCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *DeregisterAgentCommandHandler) Handle(ctx context.Context, cmd DeregisterAgentCommand) error {
	// Find agent
	a, err := h.repository.FindByID(ctx, cmd.AgentID)
	if err != nil {
		return err
	}
	if a == nil {
		return agent.ErrAgentNotFound
	}

	// Deregister agent
	a.Deregister()

	// Save agent (marks as offline)
	if err := h.repository.Save(ctx, a); err != nil {
		return err
	}

	// Publish events
	for _, event := range a.Events {
		h.eventBus.Publish(event)
	}

	return nil
}
