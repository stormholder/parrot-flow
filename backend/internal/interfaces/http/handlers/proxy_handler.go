package handlers

import (
	"context"

	proxyCommand "parrotflow/internal/application/command/proxy"
	proxyQuery "parrotflow/internal/application/query/proxy"
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

// ProxyHandler handles HTTP requests for proxy management
type ProxyHandler struct {
	createCommandHandler       *proxyCommand.CreateProxyCommandHandler
	updateCommandHandler       *proxyCommand.UpdateProxyCommandHandler
	deleteCommandHandler       *proxyCommand.DeleteProxyCommandHandler
	recordHealthCommandHandler *proxyCommand.RecordHealthCommandHandler
	activateCommandHandler     *proxyCommand.ActivateProxyCommandHandler
	deactivateCommandHandler   *proxyCommand.DeactivateProxyCommandHandler
	getQueryHandler            *proxyQuery.GetProxyQueryHandler
	listQueryHandler           *proxyQuery.ListProxiesQueryHandler
	getActiveQueryHandler      *proxyQuery.GetActiveProxiesQueryHandler
}

// NewProxyHandler creates a new ProxyHandler
func NewProxyHandler(
	createCommandHandler *proxyCommand.CreateProxyCommandHandler,
	updateCommandHandler *proxyCommand.UpdateProxyCommandHandler,
	deleteCommandHandler *proxyCommand.DeleteProxyCommandHandler,
	recordHealthCommandHandler *proxyCommand.RecordHealthCommandHandler,
	activateCommandHandler *proxyCommand.ActivateProxyCommandHandler,
	deactivateCommandHandler *proxyCommand.DeactivateProxyCommandHandler,
	getQueryHandler *proxyQuery.GetProxyQueryHandler,
	listQueryHandler *proxyQuery.ListProxiesQueryHandler,
	getActiveQueryHandler *proxyQuery.GetActiveProxiesQueryHandler,
) *ProxyHandler {
	return &ProxyHandler{
		createCommandHandler:       createCommandHandler,
		updateCommandHandler:       updateCommandHandler,
		deleteCommandHandler:       deleteCommandHandler,
		recordHealthCommandHandler: recordHealthCommandHandler,
		activateCommandHandler:     activateCommandHandler,
		deactivateCommandHandler:   deactivateCommandHandler,
		getQueryHandler:            getQueryHandler,
		listQueryHandler:           listQueryHandler,
		getActiveQueryHandler:      getActiveQueryHandler,
	}
}

// CreateProxy handles POST /api/proxies/
func (h *ProxyHandler) CreateProxy(ctx context.Context, input *commands.CreateProxyRequest) (*commands.CreateProxyResponse, error) {
	cmd := proxyCommand.CreateProxyCommand{
		Name:     input.Body.Name,
		Host:     input.Body.Host,
		Port:     input.Body.Port,
		Protocol: input.Body.Protocol,
		Username: input.Body.Username,
		Password: input.Body.Password,
		Tags:     input.Body.Tags,
	}

	p, err := h.createCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return mappers.ToCreateProxyResponse(p), nil
}

// GetProxy handles GET /api/proxies/{id}
func (h *ProxyHandler) GetProxy(ctx context.Context, input *queries.GetProxyRequest) (*queries.GetProxyResponse, error) {
	proxyID, err := proxy.NewProxyID(input.ID)
	if err != nil {
		return nil, err
	}

	query := proxyQuery.GetProxyQuery{ID: proxyID}
	p, err := h.getQueryHandler.Handle(ctx, query)
	if err != nil {
		if err == proxy.ErrProxyNotFound {
			return nil, err
		}
		return nil, err
	}

	return mappers.ToGetProxyResponse(p), nil
}

// ListProxies handles GET /api/proxies/
func (h *ProxyHandler) ListProxies(ctx context.Context, input *queries.ListProxiesRequest) (*queries.ListProxiesResponse, error) {
	query := proxyQuery.ListProxiesQuery{
		Tags: input.Tags,
	}

	// Convert empty string to nil pointer for optional status filter
	if input.Status != "" {
		query.Status = &input.Status
	}

	proxies, err := h.listQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return mappers.ToListProxiesResponse(proxies), nil
}

// GetActiveProxies handles GET /api/proxies/active
func (h *ProxyHandler) GetActiveProxies(ctx context.Context, input *queries.GetActiveProxiesRequest) (*queries.GetActiveProxiesResponse, error) {
	query := proxyQuery.GetActiveProxiesQuery{}

	proxies, err := h.getActiveQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return mappers.ToGetActiveProxiesResponse(proxies), nil
}

// UpdateProxy handles PATCH /api/proxies/{id}
func (h *ProxyHandler) UpdateProxy(ctx context.Context, input *commands.UpdateProxyRequest) (*commands.UpdateProxyResponse, error) {
	cmd := proxyCommand.UpdateProxyCommand{
		ID:       input.ID,
		Name:     input.Body.Name,
		Host:     input.Body.Host,
		Port:     input.Body.Port,
		Protocol: input.Body.Protocol,
		Username: input.Body.Username,
		Password: input.Body.Password,
	}

	p, err := h.updateCommandHandler.Handle(ctx, cmd)
	if err != nil {
		if err == proxy.ErrProxyNotFound {
			return nil, err
		}
		return nil, err
	}

	return mappers.ToUpdateProxyResponse(p), nil
}

// DeleteProxy handles DELETE /api/proxies/{id}
func (h *ProxyHandler) DeleteProxy(ctx context.Context, input *commands.DeleteProxyRequest) (*commands.DeleteProxyResponse, error) {
	cmd := proxyCommand.DeleteProxyCommand{ID: input.ID}

	err := h.deleteCommandHandler.Handle(ctx, cmd)
	if err != nil {
		if err == proxy.ErrProxyNotFound {
			return nil, err
		}
		return nil, err
	}

	response := &commands.DeleteProxyResponse{}
	response.Body.Message = "Proxy deleted successfully"
	return response, nil
}

// RecordHealth handles POST /api/proxies/{id}/health
func (h *ProxyHandler) RecordHealth(ctx context.Context, input *commands.RecordHealthRequest) (*commands.RecordHealthResponse, error) {
	cmd := proxyCommand.RecordHealthCommand{
		ID:        input.ID,
		Success:   input.Body.Success,
		LatencyMs: input.Body.LatencyMs,
		ErrorMsg:  input.Body.ErrorMsg,
	}

	p, err := h.recordHealthCommandHandler.Handle(ctx, cmd)
	if err != nil {
		if err == proxy.ErrProxyNotFound {
			return nil, err
		}
		return nil, err
	}

	return mappers.ToRecordHealthResponse(p), nil
}

// ActivateProxy handles POST /api/proxies/{id}/activate
func (h *ProxyHandler) ActivateProxy(ctx context.Context, input *commands.ActivateProxyRequest) (*commands.ActivateProxyResponse, error) {
	cmd := proxyCommand.ActivateProxyCommand{ID: input.ID}

	p, err := h.activateCommandHandler.Handle(ctx, cmd)
	if err != nil {
		if err == proxy.ErrProxyNotFound {
			return nil, err
		}
		return nil, err
	}

	return mappers.ToActivateProxyResponse(p), nil
}

// DeactivateProxy handles POST /api/proxies/{id}/deactivate
func (h *ProxyHandler) DeactivateProxy(ctx context.Context, input *commands.DeactivateProxyRequest) (*commands.DeactivateProxyResponse, error) {
	cmd := proxyCommand.DeactivateProxyCommand{ID: input.ID}

	p, err := h.deactivateCommandHandler.Handle(ctx, cmd)
	if err != nil {
		if err == proxy.ErrProxyNotFound {
			return nil, err
		}
		return nil, err
	}

	return mappers.ToDeactivateProxyResponse(p), nil
}
