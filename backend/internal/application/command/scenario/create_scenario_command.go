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

	s, err := scenario.NewScenario(scenarioID, cmd.Name)
	if err != nil {
		return nil, err
	}

	if cmd.Description != "" {
		s.UpdateDescription(cmd.Description)
	}
	if cmd.Tag != "" {
		s.UpdateTag(cmd.Tag)
	}
	if cmd.Icon != "" {
		s.UpdateIcon(cmd.Icon)
	}

	// Initialize default value objects
	// Context with single "startNode"
	startNode, err := scenario.NewNode("startNode", "start", scenario.NewPoint2D(0, 0))
	if err != nil {
		return nil, err
	}
	defaultContext := scenario.NewContext([]scenario.Node{startNode}, []scenario.Edge{})
	s.UpdateContext(defaultContext)

	// Empty InputData
	defaultInputData := scenario.NewInputData([]scenario.NodeParameters{})
	s.UpdateInputData(defaultInputData)

	// Empty Parameters with empty input/output arrays
	defaultParameters := scenario.NewParameters([]scenario.ParameterItem{}, []scenario.ParameterItem{})
	s.UpdateParameters(defaultParameters)

	if err := h.repository.Save(ctx, s); err != nil {
		return nil, err
	}

	for _, event := range s.Events {
		if err := h.eventBus.Publish(event); err != nil {
			// TODO log only or handle it somehow?
		}
	}

	s.ClearEvents()
	return s, nil
}
