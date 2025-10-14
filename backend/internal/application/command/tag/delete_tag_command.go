package command

import (
	"context"
	"errors"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
)

type DeleteTagCommand struct {
	ID tag.TagID
}

type DeleteTagCommandHandler struct {
	repository tag.Repository
	eventBus   shared.EventBus
}

func NewDeleteTagCommandHandler(repository tag.Repository, eventBus shared.EventBus) *DeleteTagCommandHandler {
	return &DeleteTagCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *DeleteTagCommandHandler) Handle(ctx context.Context, cmd DeleteTagCommand) error {
	t, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if !t.CanDelete() {
		return errors.New("cannot delete system tag")
	}

	if err := t.Delete(); err != nil {
		return err
	}

	if err := h.repository.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	for _, event := range t.Events {
		if err := h.eventBus.Publish(event); err != nil {
			// TODO: log error
		}
	}

	return nil
}
