package command

import (
	"context"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
	utils "parrotflow/pkg/shared"
)

type CreateTagCommand struct {
	Name        string
	Category    string
	Description string
	Color       string
}

type CreateTagCommandHandler struct {
	repository tag.Repository
	eventBus   shared.EventBus
}

func NewCreateTagCommandHandler(repository tag.Repository, eventBus shared.EventBus) *CreateTagCommandHandler {
	return &CreateTagCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *CreateTagCommandHandler) Handle(ctx context.Context, cmd CreateTagCommand) (*tag.Tag, error) {
	// Check if tag already exists
	exists, err := h.repository.Exists(ctx, cmd.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, tag.ErrTagAlreadyExists
	}

	tagID, err := tag.NewTagID(utils.CustomUUID())
	if err != nil {
		return nil, err
	}

	category, err := tag.NewTagCategory(cmd.Category)
	if err != nil {
		return nil, err
	}

	t, err := tag.NewTag(tagID, cmd.Name, category)
	if err != nil {
		return nil, err
	}

	if cmd.Description != "" {
		t.UpdateDescription(cmd.Description)
	}
	if cmd.Color != "" {
		if err := t.UpdateColor(cmd.Color); err != nil {
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
