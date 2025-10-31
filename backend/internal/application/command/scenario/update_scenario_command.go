package command

import (
	command "parrotflow/internal/application/command"
	"context"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
)

type UpdateScenarioCommand struct {
	ID          scenario.ScenarioID
	Name        *string
	Description *string
	Tag         *string
	Icon        *string
	Context     *scenario.Context
	InputData   *scenario.InputData
	Parameters  *scenario.Parameters
}

type UpdateScenarioCommandHandler struct {
	repository scenario.Repository
	eventBus   shared.EventBus
}

func NewUpdateScenarioCommandHandler(repository scenario.Repository, eventBus shared.EventBus) *UpdateScenarioCommandHandler {
	return &UpdateScenarioCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *UpdateScenarioCommandHandler) Handle(ctx context.Context, cmd UpdateScenarioCommand) (*scenario.Scenario, error) {
	scenario, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Name != nil {
		if err := scenario.UpdateName(*cmd.Name); err != nil {
			return nil, err
		}
	}

	if cmd.Description != nil {
		scenario.UpdateDescription(*cmd.Description)
	}

	if cmd.Tag != nil {
		scenario.UpdateTag(*cmd.Tag)
	}

	if cmd.Icon != nil {
		scenario.UpdateIcon(*cmd.Icon)
	}

	if cmd.Context != nil {
		scenario.UpdateContext(*cmd.Context)
	}

	if cmd.InputData != nil {
		scenario.UpdateInputData(*cmd.InputData)
	}

	if cmd.Parameters != nil {
		scenario.UpdateParameters(*cmd.Parameters)
	}

	// Save updated scenario
	if err := h.repository.Save(ctx, scenario); err != nil {
		return nil, err
	}

	// Publish domain events using centralized helper
	command.PublishDomainEvents(h.eventBus, scenario.Events, scenario)
	return scenario, nil
}
