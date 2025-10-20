//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeApp creates a fully wired application
func InitializeApp(db *gorm.DB) (*Application, error) {
	wire.Build(
		// Infrastructure
		NewEventBus,

		// Repositories
		RepositorySet,

		// Command Handlers
		CommandHandlerSet,

		// Query Handlers
		QueryHandlerSet,

		// HTTP Handlers
		HTTPHandlerSet,

		// Application
		NewApplication,
	)
	return nil, nil
}
