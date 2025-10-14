package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
)

type UpdateAgentCommand struct {
	AgentID      agent.AgentID
	Name         *string
	Capabilities *agent.Capabilities
	TagsToAdd    []tag.TagID
	TagsToRemove []tag.TagID
}

type UpdateAgentCommandHandler struct {
	repository agent.Repository
	eventBus   shared.EventBus
}

func NewUpdateAgentCommandHandler(
	repository agent.Repository,
	eventBus shared.EventBus,
) *UpdateAgentCommandHandler {
	return &UpdateAgentCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *UpdateAgentCommandHandler) Handle(ctx context.Context, cmd UpdateAgentCommand) (*agent.Agent, error) {
	// Find agent
	a, err := h.repository.FindByID(ctx, cmd.AgentID)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, agent.ErrAgentNotFound
	}

	// Update name if provided
	if cmd.Name != nil {
		if err := a.UpdateName(*cmd.Name); err != nil {
			return nil, err
		}
	}

	// Update capabilities if provided
	if cmd.Capabilities != nil {
		a.UpdateCapabilities(*cmd.Capabilities)
	}

	// Add tags
	for _, tagID := range cmd.TagsToAdd {
		if err := a.AddTag(tagID); err != nil {
			// Ignore if tag already exists
			continue
		}
	}

	// Remove tags
	for _, tagID := range cmd.TagsToRemove {
		a.RemoveTag(tagID)
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
