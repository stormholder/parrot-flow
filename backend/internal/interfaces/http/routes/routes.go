package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"parrotflow/internal/container"
)

// RegisterAllRoutes registers all API routes with the given application
func RegisterAllRoutes(api *huma.API, app *container.Application) {
	// System routes
	RegisterSystemRoutes(api)

	// Domain routes
	RegisterAgentRoutes(api, app.AgentHandler)
	RegisterProxyRoutes(api, app.ProxyHandler)
	RegisterTagRoutes(api, app.TagHandler)
	RegisterScenarioRoutes(api, app.ScenarioHandler)
	RegisterRunRoutes(api, app.RunHandler)
}
