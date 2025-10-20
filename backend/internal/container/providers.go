package container

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	// Domain
	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"

	// Infrastructure
	"parrotflow/internal/infrastructure/events"
	"parrotflow/internal/infrastructure/persistence"

	// Application - Commands
	agentcommand "parrotflow/internal/application/command/agent"
	proxycommand "parrotflow/internal/application/command/proxy"
	runcommand "parrotflow/internal/application/command/run"
	scenariocommand "parrotflow/internal/application/command/scenario"
	tagcommand "parrotflow/internal/application/command/tag"

	// Application - Queries
	agentquery "parrotflow/internal/application/query/agent"
	proxyquery "parrotflow/internal/application/query/proxy"
	runquery "parrotflow/internal/application/query/run"
	scenarioquery "parrotflow/internal/application/query/scenario"
	tagquery "parrotflow/internal/application/query/tag"

	// HTTP
	"parrotflow/internal/interfaces/http/handlers"
)

// ============================================================================
// INFRASTRUCTURE PROVIDERS
// ============================================================================

// NewEventBus creates a new async event bus
func NewEventBus() shared.EventBus {
	bus := events.NewAsyncEventBus()

	// Subscribe event handlers
	bus.Subscribe(events.NewScenarioCreatedHandler())
	bus.Subscribe(events.NewScenarioUpdatedHandler())
	bus.Subscribe(events.NewScenarioDeletedHandler())
	bus.Subscribe(events.NewRunCreatedHandler())
	bus.Subscribe(events.NewRunStartedHandler())
	bus.Subscribe(events.NewRunCompletedHandler())
	bus.Subscribe(events.NewRunFailedHandler())

	return bus
}

// ============================================================================
// REPOSITORY PROVIDERS
// ============================================================================

// RepositorySet provides all repositories
var RepositorySet = wire.NewSet(
	ProvideAgentRepository,
	ProvideProxyRepository,
	ProvideTagRepository,
	ProvideScenarioRepository,
	ProvideRunRepository,
)

func ProvideAgentRepository(db *gorm.DB) agent.Repository {
	return persistence.NewAgentRepository(db)
}

func ProvideProxyRepository(db *gorm.DB) proxy.Repository {
	return persistence.NewProxyRepository(db)
}

func ProvideTagRepository(db *gorm.DB) tag.Repository {
	return persistence.NewTagRepository(db)
}

func ProvideScenarioRepository(db *gorm.DB) scenario.Repository {
	return persistence.NewScenarioRepository(db)
}

func ProvideRunRepository(db *gorm.DB) run.Repository {
	return persistence.NewRunRepository(db)
}

// ============================================================================
// COMMAND HANDLER PROVIDERS
// ============================================================================

// CommandHandlerSet provides all command handlers
var CommandHandlerSet = wire.NewSet(
	// Agent commands
	agentcommand.NewRegisterAgentCommandHandler,
	agentcommand.NewUpdateHeartbeatCommandHandler,
	agentcommand.NewAssignRunCommandHandler,
	agentcommand.NewReleaseRunCommandHandler,
	agentcommand.NewUpdateAgentCommandHandler,
	agentcommand.NewDeregisterAgentCommandHandler,

	// Proxy commands
	proxycommand.NewCreateProxyCommandHandler,
	proxycommand.NewUpdateProxyCommandHandler,
	proxycommand.NewDeleteProxyCommandHandler,
	proxycommand.NewActivateProxyCommandHandler,
	proxycommand.NewDeactivateProxyCommandHandler,
	proxycommand.NewRecordHealthCommandHandler,

	// Tag commands
	tagcommand.NewCreateTagCommandHandler,
	tagcommand.NewUpdateTagCommandHandler,
	tagcommand.NewDeleteTagCommandHandler,

	// Scenario commands
	scenariocommand.NewCreateScenarioCommandHandler,
	scenariocommand.NewUpdateScenarioCommandHandler,
	scenariocommand.NewDeleteScenarioCommandHandler,

	// Run commands
	runcommand.NewCreateRunCommandHandler,
	runcommand.NewStartRunCommandHandler,
)

// ============================================================================
// QUERY HANDLER PROVIDERS
// ============================================================================

// QueryHandlerSet provides all query handlers
var QueryHandlerSet = wire.NewSet(
	// Agent queries
	agentquery.NewGetAgentQueryHandler,
	agentquery.NewListAgentsQueryHandler,
	agentquery.NewGetAvailableAgentsQueryHandler,
	agentquery.NewGetStaleAgentsQueryHandler,

	// Proxy queries
	proxyquery.NewGetProxyQueryHandler,
	proxyquery.NewListProxiesQueryHandler,
	proxyquery.NewGetActiveProxiesQueryHandler,

	// Tag queries
	tagquery.NewGetTagQueryHandler,
	tagquery.NewListTagsQueryHandler,

	// Scenario queries
	scenarioquery.NewGetScenarioQueryHandler,
	scenarioquery.NewListScenariosQueryHandler,

	// Run queries
	runquery.NewGetRunQueryHandler,
	runquery.NewListRunsQueryHandler,
)

// ============================================================================
// HTTP HANDLER PROVIDERS
// ============================================================================

// HTTPHandlerSet provides all HTTP handlers
var HTTPHandlerSet = wire.NewSet(
	handlers.NewAgentHandler,
	handlers.NewProxyHandler,
	handlers.NewTagHandler,
	handlers.NewScenarioHandler,
	handlers.NewRunHandler,
)

// ============================================================================
// APPLICATION
// ============================================================================

// Application holds all HTTP handlers
type Application struct {
	AgentHandler    *handlers.AgentHandler
	ProxyHandler    *handlers.ProxyHandler
	TagHandler      *handlers.TagHandler
	ScenarioHandler *handlers.ScenarioHandler
	RunHandler      *handlers.RunHandler
}

// NewApplication creates a new application with all dependencies wired
func NewApplication(
	agentHandler *handlers.AgentHandler,
	proxyHandler *handlers.ProxyHandler,
	tagHandler *handlers.TagHandler,
	scenarioHandler *handlers.ScenarioHandler,
	runHandler *handlers.RunHandler,
) *Application {
	return &Application{
		AgentHandler:    agentHandler,
		ProxyHandler:    proxyHandler,
		TagHandler:      tagHandler,
		ScenarioHandler: scenarioHandler,
		RunHandler:      runHandler,
	}
}
