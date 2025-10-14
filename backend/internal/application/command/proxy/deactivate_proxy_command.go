package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
	"parrotflow/pkg/shared"
)

type DeactivateProxyCommand struct {
	ID string
}

type DeactivateProxyCommandHandler struct {
	repository proxy.ProxyRepository
	eventBus   shared.EventBus
}

func NewDeactivateProxyCommandHandler(repository proxy.ProxyRepository, eventBus shared.EventBus) *DeactivateProxyCommandHandler {
	return &DeactivateProxyCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *DeactivateProxyCommandHandler) Handle(ctx context.Context, cmd DeactivateProxyCommand) (*proxy.Proxy, error) {
	// Parse proxy ID
	proxyID, err := proxy.NewProxyID(cmd.ID)
	if err != nil {
		return nil, err
	}

	// Find proxy
	p, err := h.repository.FindByID(ctx, proxyID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, proxy.ErrProxyNotFound
	}

	// Deactivate proxy
	if err := p.Deactivate(); err != nil {
		return nil, err
	}

	// Save to repository
	if err := h.repository.Save(ctx, p); err != nil {
		return nil, err
	}

	// Publish domain events
	for _, event := range p.Events {
		h.eventBus.Publish(event)
	}

	return p, nil
}
