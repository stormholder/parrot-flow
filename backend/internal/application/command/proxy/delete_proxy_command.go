package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/domain/shared"
)

type DeleteProxyCommand struct {
	ID string
}

type DeleteProxyCommandHandler struct {
	repository proxy.Repository
	eventBus   shared.EventBus
}

func NewDeleteProxyCommandHandler(repository proxy.Repository, eventBus shared.EventBus) *DeleteProxyCommandHandler {
	return &DeleteProxyCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *DeleteProxyCommandHandler) Handle(ctx context.Context, cmd DeleteProxyCommand) error {
	// Parse proxy ID
	proxyID, err := proxy.NewProxyID(cmd.ID)
	if err != nil {
		return err
	}

	// Check if proxy exists
	p, err := h.repository.FindByID(ctx, proxyID)
	if err != nil {
		return err
	}
	if p == nil {
		return proxy.ErrProxyNotFound
	}

	// Delete from repository
	if err := h.repository.Delete(ctx, proxyID); err != nil {
		return err
	}

	// Publish domain events (if any were generated before deletion)
	for _, event := range p.Events {
		h.eventBus.Publish(event)
	}

	return nil
}
