package routes

import (
	agentcommand "parrotflow/internal/application/command/agent"
	agentquery "parrotflow/internal/application/query/agent"
	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/interfaces/http/handlers"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAgentRoutes(
	api *huma.API,
	agentRepository agent.Repository,
	eventBus shared.EventBus,
) {
	// Create command handlers
	registerCommandHandler := agentcommand.NewRegisterAgentCommandHandler(agentRepository, eventBus)
	updateHeartbeatCommandHandler := agentcommand.NewUpdateHeartbeatCommandHandler(agentRepository, eventBus)
	assignRunCommandHandler := agentcommand.NewAssignRunCommandHandler(agentRepository, eventBus)
	releaseRunCommandHandler := agentcommand.NewReleaseRunCommandHandler(agentRepository, eventBus)
	updateCommandHandler := agentcommand.NewUpdateAgentCommandHandler(agentRepository, eventBus)
	deregisterCommandHandler := agentcommand.NewDeregisterAgentCommandHandler(agentRepository, eventBus)

	// Create query handlers
	getQueryHandler := agentquery.NewGetAgentQueryHandler(agentRepository)
	listQueryHandler := agentquery.NewListAgentsQueryHandler(agentRepository)
	getAvailableQueryHandler := agentquery.NewGetAvailableAgentsQueryHandler(agentRepository)
	getStaleQueryHandler := agentquery.NewGetStaleAgentsQueryHandler(agentRepository)

	// Create handler
	handler := handlers.NewAgentHandler(
		registerCommandHandler,
		updateHeartbeatCommandHandler,
		assignRunCommandHandler,
		releaseRunCommandHandler,
		updateCommandHandler,
		deregisterCommandHandler,
		getQueryHandler,
		listQueryHandler,
		getAvailableQueryHandler,
		getStaleQueryHandler,
	)

	// Register agent - POST /api/agents/
	huma.Register(*api, huma.Operation{
		OperationID: "register-agent",
		Method:      "POST",
		Path:        "/api/agents/",
		Summary:     "Register a new agent",
		Description: "Registers a new agent with the system (auto-registration on first heartbeat)",
		Tags:        []string{"agents"},
	}, handler.RegisterAgent)

	// Update heartbeat - POST /api/agents/{id}/heartbeat
	huma.Register(*api, huma.Operation{
		OperationID: "update-agent-heartbeat",
		Method:      "POST",
		Path:        "/api/agents/{id}/heartbeat",
		Summary:     "Update agent heartbeat",
		Description: "Updates the agent's last heartbeat timestamp and auto-recovers from disconnected status",
		Tags:        []string{"agents"},
	}, handler.UpdateHeartbeat)

	// Assign run - POST /api/agents/{id}/assign-run
	huma.Register(*api, huma.Operation{
		OperationID: "assign-run-to-agent",
		Method:      "POST",
		Path:        "/api/agents/{id}/assign-run",
		Summary:     "Assign a run to an agent",
		Description: "Assigns a run to the agent, incrementing its run count and updating status",
		Tags:        []string{"agents"},
	}, handler.AssignRun)

	// Release run - POST /api/agents/{id}/release-run
	huma.Register(*api, huma.Operation{
		OperationID: "release-run-from-agent",
		Method:      "POST",
		Path:        "/api/agents/{id}/release-run",
		Summary:     "Release a run from an agent",
		Description: "Releases a run from the agent, decrementing its run count and updating status",
		Tags:        []string{"agents"},
	}, handler.ReleaseRun)

	// Get agent - GET /api/agents/{id}
	huma.Register(*api, huma.Operation{
		OperationID: "get-agent",
		Method:      "GET",
		Path:        "/api/agents/{id}",
		Summary:     "Get an agent by ID",
		Description: "Retrieves a single agent by its ID",
		Tags:        []string{"agents"},
	}, handler.GetAgent)

	// List agents - GET /api/agents/
	huma.Register(*api, huma.Operation{
		OperationID: "list-agents",
		Method:      "GET",
		Path:        "/api/agents/",
		Summary:     "List agents",
		Description: "Lists all agents with optional filters (status, browser_type, platform, tags, only_healthy)",
		Tags:        []string{"agents"},
	}, handler.ListAgents)

	// Get available agents - GET /api/agents/available
	huma.Register(*api, huma.Operation{
		OperationID: "get-available-agents",
		Method:      "GET",
		Path:        "/api/agents/available",
		Summary:     "Get available agents",
		Description: "Retrieves agents that can accept new runs, with optional filters (browser_type, platform, tags)",
		Tags:        []string{"agents"},
	}, handler.GetAvailableAgents)

	// Get stale agents - GET /api/agents/stale
	huma.Register(*api, huma.Operation{
		OperationID: "get-stale-agents",
		Method:      "GET",
		Path:        "/api/agents/stale",
		Summary:     "Get stale agents",
		Description: "Retrieves agents that haven't sent a heartbeat within the timeout period",
		Tags:        []string{"agents"},
	}, handler.GetStaleAgents)

	// Update agent - PATCH /api/agents/{id}
	huma.Register(*api, huma.Operation{
		OperationID: "update-agent",
		Method:      "PATCH",
		Path:        "/api/agents/{id}",
		Summary:     "Update an agent",
		Description: "Updates agent properties (name, capabilities, tags)",
		Tags:        []string{"agents"},
	}, handler.UpdateAgent)

	// Deregister agent - DELETE /api/agents/{id}
	huma.Register(*api, huma.Operation{
		OperationID: "deregister-agent",
		Method:      "DELETE",
		Path:        "/api/agents/{id}",
		Summary:     "Deregister an agent",
		Description: "Marks the agent as offline and deregisters it from the system",
		Tags:        []string{"agents"},
	}, handler.DeregisterAgent)
}
