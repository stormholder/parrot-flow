package routes

import (
	"github.com/danielgtaylor/huma/v2"

	proxyCommand "parrotflow/internal/application/command/proxy"
	proxyQuery "parrotflow/internal/application/query/proxy"
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/interfaces/http/handlers"
)

// RegisterProxyRoutes registers all proxy-related routes
func RegisterProxyRoutes(api *huma.API, proxyRepository proxy.Repository, eventBus shared.EventBus) {
	// Initialize command handlers
	createCommandHandler := proxyCommand.NewCreateProxyCommandHandler(proxyRepository, eventBus)
	updateCommandHandler := proxyCommand.NewUpdateProxyCommandHandler(proxyRepository, eventBus)
	deleteCommandHandler := proxyCommand.NewDeleteProxyCommandHandler(proxyRepository, eventBus)
	recordHealthCommandHandler := proxyCommand.NewRecordHealthCommandHandler(proxyRepository, eventBus)
	activateCommandHandler := proxyCommand.NewActivateProxyCommandHandler(proxyRepository, eventBus)
	deactivateCommandHandler := proxyCommand.NewDeactivateProxyCommandHandler(proxyRepository, eventBus)

	// Initialize query handlers
	getQueryHandler := proxyQuery.NewGetProxyQueryHandler(proxyRepository)
	listQueryHandler := proxyQuery.NewListProxiesQueryHandler(proxyRepository)
	getActiveQueryHandler := proxyQuery.NewGetActiveProxiesQueryHandler(proxyRepository)

	// Initialize HTTP handler
	handler := handlers.NewProxyHandler(
		createCommandHandler,
		updateCommandHandler,
		deleteCommandHandler,
		recordHealthCommandHandler,
		activateCommandHandler,
		deactivateCommandHandler,
		getQueryHandler,
		listQueryHandler,
		getActiveQueryHandler,
	)

	// Register routes

	// POST /api/proxies/ - Create a new proxy
	huma.Register(*api, huma.Operation{
		OperationID: "create-proxy",
		Method:      "POST",
		Path:        "/api/proxies/",
		Summary:     "Create a new proxy",
		Description: "Creates a new proxy server configuration with optional authentication credentials and tags",
		Tags:        []string{"proxies"},
	}, handler.CreateProxy)

	// GET /api/proxies/{id} - Get a single proxy by ID
	huma.Register(*api, huma.Operation{
		OperationID: "get-proxy",
		Method:      "GET",
		Path:        "/api/proxies/{id}",
		Summary:     "Get proxy by ID",
		Description: "Retrieves a single proxy configuration by its unique identifier",
		Tags:        []string{"proxies"},
	}, handler.GetProxy)

	// GET /api/proxies/ - List all proxies with optional filters
	huma.Register(*api, huma.Operation{
		OperationID: "list-proxies",
		Method:      "GET",
		Path:        "/api/proxies/",
		Summary:     "List proxies",
		Description: "Lists all proxies with optional filtering by status or tags",
		Tags:        []string{"proxies"},
	}, handler.ListProxies)

	// GET /api/proxies/active - Get all active proxies
	huma.Register(*api, huma.Operation{
		OperationID: "get-active-proxies",
		Method:      "GET",
		Path:        "/api/proxies/active",
		Summary:     "Get active proxies",
		Description: "Retrieves all proxies with active status",
		Tags:        []string{"proxies"},
	}, handler.GetActiveProxies)

	// PATCH /api/proxies/{id} - Update an existing proxy
	huma.Register(*api, huma.Operation{
		OperationID: "update-proxy",
		Method:      "PATCH",
		Path:        "/api/proxies/{id}",
		Summary:     "Update proxy",
		Description: "Updates an existing proxy's configuration (name, host, port, protocol, credentials)",
		Tags:        []string{"proxies"},
	}, handler.UpdateProxy)

	// DELETE /api/proxies/{id} - Delete a proxy
	huma.Register(*api, huma.Operation{
		OperationID: "delete-proxy",
		Method:      "DELETE",
		Path:        "/api/proxies/{id}",
		Summary:     "Delete proxy",
		Description: "Deletes a proxy configuration",
		Tags:        []string{"proxies"},
	}, handler.DeleteProxy)

	// POST /api/proxies/{id}/health - Record health check result
	huma.Register(*api, huma.Operation{
		OperationID: "record-proxy-health",
		Method:      "POST",
		Path:        "/api/proxies/{id}/health",
		Summary:     "Record health check",
		Description: "Records the result of a health check (success or failure) for a proxy",
		Tags:        []string{"proxies"},
	}, handler.RecordHealth)

	// POST /api/proxies/{id}/activate - Activate a proxy
	huma.Register(*api, huma.Operation{
		OperationID: "activate-proxy",
		Method:      "POST",
		Path:        "/api/proxies/{id}/activate",
		Summary:     "Activate proxy",
		Description: "Activates a proxy, making it available for use",
		Tags:        []string{"proxies"},
	}, handler.ActivateProxy)

	// POST /api/proxies/{id}/deactivate - Deactivate a proxy
	huma.Register(*api, huma.Operation{
		OperationID: "deactivate-proxy",
		Method:      "POST",
		Path:        "/api/proxies/{id}/deactivate",
		Summary:     "Deactivate proxy",
		Description: "Deactivates a proxy, removing it from the available pool",
		Tags:        []string{"proxies"},
	}, handler.DeactivateProxy)
}
