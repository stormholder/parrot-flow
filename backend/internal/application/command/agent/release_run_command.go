package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
)

type ReleaseRunCommand struct {
	AgentID agent.AgentID
}

type ReleaseRunCommandHandler struct {
	repository agent.Repository
	eventBus   shared.EventBus
}

func NewReleaseRunCommandHandler(
	repository agent.Repository,
	eventBus shared.EventBus,
) *ReleaseRunCommandHandler {
	return &ReleaseRunCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *ReleaseRunCommandHandler) Handle(ctx context.Context, cmd ReleaseRunCommand) (*agent.Agent, error) {
	// Find agent
	a, err := h.repository.FindByID(ctx, cmd.AgentID)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, agent.ErrAgentNotFound
	}

	// Release run
	if err := a.ReleaseRun(); err != nil {
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
