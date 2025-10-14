package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
	utils "parrotflow/pkg/shared"
)

type RegisterAgentCommand struct {
	Name           string
	Capabilities   agent.Capabilities
	ConnectionInfo agent.ConnectionInfo
}

type RegisterAgentCommandHandler struct {
	repository agent.Repository
	eventBus   shared.EventBus
}

func NewRegisterAgentCommandHandler(
	repository agent.Repository,
	eventBus shared.EventBus,
) *RegisterAgentCommandHandler {
	return &RegisterAgentCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *RegisterAgentCommandHandler) Handle(ctx context.Context, cmd RegisterAgentCommand) (*agent.Agent, error) {
	// Check if agent with the same name already exists
	exists, err := h.repository.ExistsByName(ctx, cmd.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, agent.ErrAgentAlreadyExists
	}

	// Create agent ID
	agentID, err := agent.NewAgentID(utils.CustomUUID())
	if err != nil {
		return nil, err
	}

	// Create agent (auto-registration)
	a, err := agent.NewAgent(agentID, cmd.Name, cmd.Capabilities, cmd.ConnectionInfo)
	if err != nil {
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
