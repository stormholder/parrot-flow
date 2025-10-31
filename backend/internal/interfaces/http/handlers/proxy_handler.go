package handlers

import (
	"context"

	command "parrotflow/internal/application/command/proxy"
	query "parrotflow/internal/application/query/proxy"
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

type ProxyHandler struct {
	// Commands
	createCommandHandler     *command.CreateProxyCommandHandler
	updateCommandHandler     *command.UpdateProxyCommandHandler
	deleteCommandHandler     *command.DeleteProxyCommandHandler
	activateCommandHandler   *command.ActivateProxyCommandHandler
	deactivateCommandHandler *command.DeactivateProxyCommandHandler
	recordHealthCommandHandler *command.RecordHealthCommandHandler

	// Queries
	getQueryHandler    *query.GetProxyQueryHandler
	listQueryHandler   *query.ListProxiesQueryHandler
	activeQueryHandler *query.GetActiveProxiesQueryHandler

	// Mappers - using functional types
	createMapper       mappers.CreateMapperFunc[*proxy.Proxy, *commands.CreateProxyResponse]
	updateMapper       mappers.UpdateMapperFunc[*proxy.Proxy, *commands.UpdateProxyResponse]
	deleteMapper       mappers.DeleteMapperFunc[*commands.DeleteProxyResponse]
	activateMapper     mappers.CreateMapperFunc[*proxy.Proxy, *commands.ActivateProxyResponse]
	deactivateMapper   mappers.CreateMapperFunc[*proxy.Proxy, *commands.DeactivateProxyResponse]
	recordHealthMapper mappers.CreateMapperFunc[*proxy.Proxy, *commands.RecordHealthResponse]
	getMapper          mappers.GetMapperFunc[*proxy.Proxy, *queries.GetProxyResponse]
	listMapper         mappers.ListMapperFunc[proxy.Proxy, *queries.ListProxiesResponse]
	activeListMapper   mappers.ListMapperFunc[proxy.Proxy, *queries.GetActiveProxiesResponse]
}

func NewProxyHandler(
	createCommandHandler *command.CreateProxyCommandHandler,
	updateCommandHandler *command.UpdateProxyCommandHandler,
	deleteCommandHandler *command.DeleteProxyCommandHandler,
	recordHealthCommandHandler *command.RecordHealthCommandHandler,
	activateCommandHandler *command.ActivateProxyCommandHandler,
	deactivateCommandHandler *command.DeactivateProxyCommandHandler,
	getQueryHandler *query.GetProxyQueryHandler,
	listQueryHandler *query.ListProxiesQueryHandler,
	activeQueryHandler *query.GetActiveProxiesQueryHandler,
) *ProxyHandler {
	return &ProxyHandler{
		createCommandHandler:       createCommandHandler,
		updateCommandHandler:       updateCommandHandler,
		deleteCommandHandler:       deleteCommandHandler,
		activateCommandHandler:     activateCommandHandler,
		deactivateCommandHandler:   deactivateCommandHandler,
		recordHealthCommandHandler: recordHealthCommandHandler,
		getQueryHandler:            getQueryHandler,
		listQueryHandler:           listQueryHandler,
		activeQueryHandler:         activeQueryHandler,
		createMapper:               mappers.ProxyCreateMapper,
		updateMapper:               mappers.ProxyUpdateMapper,
		deleteMapper:               mappers.ProxyDeleteMapper,
		activateMapper:             mappers.ProxyActivateMapper,
		deactivateMapper:           mappers.ProxyDeactivateMapper,
		recordHealthMapper:         mappers.ProxyRecordHealthMapper,
		getMapper:                  mappers.ProxyGetMapper,
		listMapper:                 mappers.ProxyListMapper,
		activeListMapper:           mappers.ProxyActiveListMapper,
	}
}

func (h *ProxyHandler) CreateProxy(ctx context.Context, req *commands.CreateProxyRequest) (*commands.CreateProxyResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.CreateProxyRequest) (command.CreateProxyCommand, error) {
			return command.CreateProxyCommand{
				Name:     r.Body.Name,
				Host:     r.Body.Host,
				Port:     r.Body.Port,
				Protocol: r.Body.Protocol,
				Username: r.Body.Username,
				Password: r.Body.Password,
			}, nil
		},
		CommandHandlerFunc[command.CreateProxyCommand, *proxy.Proxy](h.createCommandHandler.Handle),
		h.createMapper,
	)
}

func (h *ProxyHandler) UpdateProxy(ctx context.Context, req *commands.UpdateProxyRequest) (*commands.UpdateProxyResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.UpdateProxyRequest) (command.UpdateProxyCommand, error) {
			return command.UpdateProxyCommand{
				ID:       r.ID,
				Name:     r.Body.Name,
				Host:     r.Body.Host,
				Port:     r.Body.Port,
				Protocol: r.Body.Protocol,
			}, nil
		},
		CommandHandlerFunc[command.UpdateProxyCommand, *proxy.Proxy](h.updateCommandHandler.Handle),
		h.updateMapper,
	)
}

func (h *ProxyHandler) DeleteProxy(ctx context.Context, req *commands.DeleteProxyRequest) (*commands.DeleteProxyResponse, error) {
	return HandleSimpleCommand(
		ctx,
		req,
		func(r *commands.DeleteProxyRequest) (command.DeleteProxyCommand, error) {
			return command.DeleteProxyCommand{ID: r.ID}, nil
		},
		SimpleCommandHandlerFunc[command.DeleteProxyCommand](h.deleteCommandHandler.Handle),
		h.deleteMapper.Map,
	)
}

func (h *ProxyHandler) ActivateProxy(ctx context.Context, req *commands.ActivateProxyRequest) (*commands.ActivateProxyResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.ActivateProxyRequest) (command.ActivateProxyCommand, error) {
			return command.ActivateProxyCommand{ID: r.ID}, nil
		},
		CommandHandlerFunc[command.ActivateProxyCommand, *proxy.Proxy](h.activateCommandHandler.Handle),
		h.activateMapper,
	)
}

func (h *ProxyHandler) DeactivateProxy(ctx context.Context, req *commands.DeactivateProxyRequest) (*commands.DeactivateProxyResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.DeactivateProxyRequest) (command.DeactivateProxyCommand, error) {
			return command.DeactivateProxyCommand{ID: r.ID}, nil
		},
		CommandHandlerFunc[command.DeactivateProxyCommand, *proxy.Proxy](h.deactivateCommandHandler.Handle),
		h.deactivateMapper,
	)
}

func (h *ProxyHandler) RecordHealth(ctx context.Context, req *commands.RecordHealthRequest) (*commands.RecordHealthResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.RecordHealthRequest) (command.RecordHealthCommand, error) {
			return command.RecordHealthCommand{
				ID:        r.ID,
				Success:   r.Body.Success,
				LatencyMs: r.Body.LatencyMs,
				ErrorMsg:  r.Body.ErrorMsg,
			}, nil
		},
		CommandHandlerFunc[command.RecordHealthCommand, *proxy.Proxy](h.recordHealthCommandHandler.Handle),
		h.recordHealthMapper,
	)
}

func (h *ProxyHandler) GetProxy(ctx context.Context, req *queries.GetProxyRequest) (*queries.GetProxyResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetProxyRequest) (query.GetProxyQuery, error) {
			proxyID, err := proxy.NewProxyID(r.ID)
			if err != nil {
				return query.GetProxyQuery{}, err
			}
			return query.GetProxyQuery{ID: proxyID}, nil
		},
		QueryHandlerFunc[query.GetProxyQuery, *proxy.Proxy](h.getQueryHandler.Handle),
		h.getMapper,
	)
}

func (h *ProxyHandler) ListProxies(ctx context.Context, req *queries.ListProxiesRequest) (*queries.ListProxiesResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.ListProxiesRequest) (query.ListProxiesQuery, error) {
			return query.ListProxiesQuery{}, nil
		},
		QueryHandlerFunc[query.ListProxiesQuery, []*proxy.Proxy](h.listQueryHandler.Handle),
		h.listMapper,
	)
}

func (h *ProxyHandler) GetActiveProxies(ctx context.Context, req *queries.GetActiveProxiesRequest) (*queries.GetActiveProxiesResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetActiveProxiesRequest) (query.GetActiveProxiesQuery, error) {
			return query.GetActiveProxiesQuery{}, nil
		},
		QueryHandlerFunc[query.GetActiveProxiesQuery, []*proxy.Proxy](h.activeQueryHandler.Handle),
		h.activeListMapper,
	)
}
