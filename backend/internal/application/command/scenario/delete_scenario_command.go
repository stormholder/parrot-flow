package command

import (
	"context"
	"errors"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
)

type DeleteScenarioCommand struct {
	ID scenario.ScenarioID
}

type DeleteScenarioCommandHandler struct {
	repository scenario.Repository
	eventBus   shared.EventBus
}

func NewDeleteScenarioCommandHandler(repository scenario.Repository, eventBus shared.EventBus) *DeleteScenarioCommandHandler {
	return &DeleteScenarioCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *DeleteScenarioCommandHandler) Handle(ctx context.Context, cmd DeleteScenarioCommand) error {
	exists, err := h.repository.Exists(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("scenario not found")
	}

	if err := h.repository.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	event := scenario.ScenarioDeleted{
		BaseEvent:  shared.NewBaseEvent("ScenarioDeleted", cmd.ID.String()),
		ScenarioID: cmd.ID.String(),
	}

	if err := h.eventBus.Publish(event); err != nil {
		// TODO
	}

	return nil
}
