package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
	"parrotflow/pkg/shared"
	"parrotflow/pkg/shared/utils"
)

type CreateProxyCommand struct {
	Name     string
	Host     string
	Port     int
	Protocol string
	Username string
	Password string
	Tags     []string // Tag IDs
}

type CreateProxyCommandHandler struct {
	repository proxy.ProxyRepository
	eventBus   shared.EventBus
}

func NewCreateProxyCommandHandler(repository proxy.ProxyRepository, eventBus shared.EventBus) *CreateProxyCommandHandler {
	return &CreateProxyCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *CreateProxyCommandHandler) Handle(ctx context.Context, cmd CreateProxyCommand) (*proxy.Proxy, error) {
	// Check if proxy already exists
	exists, err := h.repository.Exists(ctx, cmd.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, proxy.ErrProxyAlreadyExists
	}

	// Create proxy ID
	proxyID, err := proxy.NewProxyID(utils.CustomUUID())
	if err != nil {
		return nil, err
	}

	// Create protocol
	protocol, err := proxy.NewProxyProtocol(cmd.Protocol)
	if err != nil {
		return nil, err
	}

	// Create proxy
	p, err := proxy.NewProxy(proxyID, cmd.Name, cmd.Host, cmd.Port, protocol)
	if err != nil {
		return nil, err
	}

	// Set credentials if provided
	if cmd.Username != "" || cmd.Password != "" {
		credentials, err := proxy.NewProxyCredentials(cmd.Username, cmd.Password)
		if err != nil {
			return nil, err
		}
		if err := p.SetCredentials(credentials); err != nil {
			return nil, err
		}
	}

	// Add tags if provided
	if len(cmd.Tags) > 0 {
		for _, tagIDStr := range cmd.Tags {
			// Note: Tag validation should be done at the API layer
			p.AddTag(tagIDStr)
		}
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
