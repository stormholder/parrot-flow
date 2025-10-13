package proxy

import (
	"context"
	"parrotflow/internal/domain/tag"
)

// Repository defines the interface for proxy persistence
type Repository interface {
	// Save persists a proxy
	Save(ctx context.Context, proxy *Proxy) error

	// FindByID retrieves a proxy by its ID
	FindByID(ctx context.Context, id ProxyID) (*Proxy, error)

	// FindByName retrieves a proxy by its name
	FindByName(ctx context.Context, name string) (*Proxy, error)

	// FindAll retrieves all proxies
	FindAll(ctx context.Context) ([]*Proxy, error)

	// FindByStatus retrieves proxies by status
	FindByStatus(ctx context.Context, status ProxyStatus) ([]*Proxy, error)

	// FindByTags retrieves proxies that have all specified tags
	FindByTags(ctx context.Context, tagIDs []tag.TagID) ([]*Proxy, error)

	// FindActive retrieves all active proxies
	FindActive(ctx context.Context) ([]*Proxy, error)

	// Delete removes a proxy
	Delete(ctx context.Context, id ProxyID) error

	// Exists checks if a proxy with the given name already exists
	Exists(ctx context.Context, name string) (bool, error)
}
