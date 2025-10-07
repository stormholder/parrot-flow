package command

import (
	"context"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/shared"
)

type StartRunCommand struct {
	RunID run.RunID
}

type StartRunCommandHandler struct {
	repository run.Repository
	eventBus   shared.EventBus
}

func NewStartRunCommandHandler(repository run.Repository, eventBus shared.EventBus) *StartRunCommandHandler {
	return &StartRunCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *StartRunCommandHandler) Handle(ctx context.Context, cmd StartRunCommand) (*run.Run, error) {
	run, err := h.repository.FindByID(ctx, cmd.RunID)
	if err != nil {
		return nil, err
	}

	if err := run.Start(); err != nil {
		return nil, err
	}

	if err := h.repository.Save(ctx, run); err != nil {
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
