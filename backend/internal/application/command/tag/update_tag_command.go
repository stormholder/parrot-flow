package command

import (
	"context"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
)

type UpdateTagCommand struct {
	ID          tag.TagID
	Description *string
	Color       *string
}

type UpdateTagCommandHandler struct {
	repository tag.Repository
	eventBus   shared.EventBus
}

func NewUpdateTagCommandHandler(repository tag.Repository, eventBus shared.EventBus) *UpdateTagCommandHandler {
	return &UpdateTagCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *UpdateTagCommandHandler) Handle(ctx context.Context, cmd UpdateTagCommand) (*tag.Tag, error) {
	t, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Description != nil {
		t.UpdateDescription(*cmd.Description)
	}

	if cmd.Color != nil {
		if err := t.UpdateColor(*cmd.Color); err != nil {
			return nil, err
		}
	}

	if err := h.repository.Save(ctx, t); err != nil {
		return nil, err
	}

	for _, event := range t.Events {
		if err := h.eventBus.Publish(event); err != nil {
			// TODO: log error
		}
	}

	t.ClearEvents()
	return t, nil
}
