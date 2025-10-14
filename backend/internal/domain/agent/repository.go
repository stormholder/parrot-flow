package agent

import (
	"context"
	"parrotflow/internal/domain/tag"
	"time"
)

// Repository defines the interface for agent persistence
type Repository interface {
	// Save persists an agent
	Save(ctx context.Context, agent *Agent) error

	// FindByID retrieves an agent by its ID
	FindByID(ctx context.Context, id AgentID) (*Agent, error)

	// FindByName retrieves an agent by its name
	FindByName(ctx context.Context, name string) (*Agent, error)

	// FindAll retrieves all agents
	FindAll(ctx context.Context) ([]*Agent, error)

	// FindByStatus retrieves agents by status
	FindByStatus(ctx context.Context, status AgentStatus) ([]*Agent, error)

	// FindByTags retrieves agents that have all specified tags
	FindByTags(ctx context.Context, tagIDs []tag.TagID) ([]*Agent, error)

	// FindAvailable retrieves agents that can accept new runs
	FindAvailable(ctx context.Context) ([]*Agent, error)

	// FindByCapability retrieves agents that support a specific browser type
	FindByBrowserType(ctx context.Context, browserType BrowserType) ([]*Agent, error)

	// FindByPlatform retrieves agents running on a specific platform
	FindByPlatform(ctx context.Context, platform Platform) ([]*Agent, error)

	// FindStaleAgents retrieves agents whose last heartbeat is older than the timeout
	FindStaleAgents(ctx context.Context, heartbeatTimeout time.Duration) ([]*Agent, error)

	// Delete removes an agent
	Delete(ctx context.Context, id AgentID) error

	// Exists checks if an agent with the given ID exists
	Exists(ctx context.Context, id AgentID) (bool, error)

	// ExistsByName checks if an agent with the given name already exists
	ExistsByName(ctx context.Context, name string) (bool, error)
}
