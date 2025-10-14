package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
	"parrotflow/pkg/shared"
)

type UpdateProxyCommand struct {
	ID       string
	Name     *string
	Host     *string
	Port     *int
	Protocol *string
	Username *string
	Password *string
}

type UpdateProxyCommandHandler struct {
	repository proxy.ProxyRepository
	eventBus   shared.EventBus
}

func NewUpdateProxyCommandHandler(repository proxy.ProxyRepository, eventBus shared.EventBus) *UpdateProxyCommandHandler {
	return &UpdateProxyCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *UpdateProxyCommandHandler) Handle(ctx context.Context, cmd UpdateProxyCommand) (*proxy.Proxy, error) {
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

	// Update name if provided
	if cmd.Name != nil && *cmd.Name != "" {
		p.Name = *cmd.Name
	}

	// Update host if provided
	if cmd.Host != nil && *cmd.Host != "" {
		p.Host = *cmd.Host
	}

	// Update port if provided
	if cmd.Port != nil && *cmd.Port > 0 {
		p.Port = *cmd.Port
	}

	// Update protocol if provided
	if cmd.Protocol != nil && *cmd.Protocol != "" {
		protocol, err := proxy.NewProxyProtocol(*cmd.Protocol)
		if err != nil {
			return nil, err
		}
		p.Protocol = protocol
	}

	// Update credentials if provided
	if cmd.Username != nil || cmd.Password != nil {
		username := ""
		password := ""
		if cmd.Username != nil {
			username = *cmd.Username
		} else if p.Credentials != nil {
			username = p.Credentials.Username
		}
		if cmd.Password != nil {
			password = *cmd.Password
		} else if p.Credentials != nil {
			password = p.Credentials.Password
		}

		if username != "" || password != "" {
			credentials, err := proxy.NewProxyCredentials(username, password)
			if err != nil {
				return nil, err
			}
			if err := p.SetCredentials(credentials); err != nil {
				return nil, err
			}
		}
	}

	// Mark as updated
	p.UpdatedAt = shared.TimestampNow()

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
