package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
	"parrotflow/pkg/shared"
)

type ActivateProxyCommand struct {
	ID string
}

type ActivateProxyCommandHandler struct {
	repository proxy.ProxyRepository
	eventBus   shared.EventBus
}

func NewActivateProxyCommandHandler(repository proxy.ProxyRepository, eventBus shared.EventBus) *ActivateProxyCommandHandler {
	return &ActivateProxyCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *ActivateProxyCommandHandler) Handle(ctx context.Context, cmd ActivateProxyCommand) (*proxy.Proxy, error) {
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

	// Activate proxy
	if err := p.Activate(); err != nil {
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
