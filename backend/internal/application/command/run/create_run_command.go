package command

import (
	"context"
	"errors"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
	utils "parrotflow/pkg/shared"
)

type CreateRunCommand struct {
	ScenarioID scenario.ScenarioID
	Parameters string
}

type CreateRunCommandHandler struct {
	runRepository      run.Repository
	scenarioRepository scenario.Repository
	eventBus           shared.EventBus
}

func NewCreateRunCommandHandler(
	runRepository run.Repository,
	scenarioRepository scenario.Repository,
	eventBus shared.EventBus,
) *CreateRunCommandHandler {
	return &CreateRunCommandHandler{
		runRepository:      runRepository,
		scenarioRepository: scenarioRepository,
		eventBus:           eventBus,
	}
}

func (h *CreateRunCommandHandler) Handle(ctx context.Context, cmd CreateRunCommand) (*run.Run, error) {
	exists, err := h.scenarioRepository.Exists(ctx, cmd.ScenarioID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("scenario not found")
	}

	runID, err := run.NewRunID(utils.CustomUUID())
	if err != nil {
		return nil, err
	}

	run, err := run.NewRun(runID, cmd.ScenarioID, cmd.Parameters)
	if err != nil {
		return nil, err
	}

	if err := h.runRepository.Save(ctx, run); err != nil {
		return nil, err
	}

	for _, event := range run.Events {
		if err := h.eventBus.Publish(event); err != nil {
			// TODO log
		}
	}

	run.ClearEvents()
	return run, nil
}
