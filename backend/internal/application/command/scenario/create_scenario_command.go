package command

import (
	"context"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
	utils "parrotflow/pkg/shared"
)

type CreateScenarioCommand struct {
	Name        string
	Description string
	Tag         string
	Icon        string
}

type CreateScenarioCommandHandler struct {
	repository scenario.Repository
	eventBus   shared.EventBus
}

func NewCreateScenarioCommandHandler(repository scenario.Repository, eventBus shared.EventBus) *CreateScenarioCommandHandler {
	return &CreateScenarioCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *CreateScenarioCommandHandler) Handle(ctx context.Context, cmd CreateScenarioCommand) (*scenario.Scenario, error) {
	scenarioID, err := scenario.NewScenarioID(utils.CustomUUID())
	if err != nil {
		return nil, err
	}

	scenario, err := scenario.NewScenario(scenarioID, cmd.Name)
	if err != nil {
		return nil, err
	}

	if cmd.Description != "" {
		scenario.UpdateDescription(cmd.Description)
	}
	if cmd.Tag != "" {
		scenario.UpdateTag(cmd.Tag)
	}
	if cmd.Icon != "" {
		scenario.UpdateIcon(cmd.Icon)
	}

	if err := h.repository.Save(ctx, scenario); err != nil {
		return nil, err
	}

	for _, event := range scenario.Events {
		if err := h.eventBus.Publish(event); err != nil {
			// TODO log only or handle it somehow?
		}
	}

	scenario.ClearEvents()
	return scenario, nil
}
