package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
)

type AssignRunCommand struct {
	AgentID agent.AgentID
}

type AssignRunCommandHandler struct {
	repository agent.Repository
	eventBus   shared.EventBus
}

func NewAssignRunCommandHandler(
	repository agent.Repository,
	eventBus shared.EventBus,
) *AssignRunCommandHandler {
	return &AssignRunCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *AssignRunCommandHandler) Handle(ctx context.Context, cmd AssignRunCommand) (*agent.Agent, error) {
	// Find agent
	a, err := h.repository.FindByID(ctx, cmd.AgentID)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, agent.ErrAgentNotFound
	}

	// Assign run
	if err := a.AssignRun(); err != nil {
		return nil, err
	}

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
