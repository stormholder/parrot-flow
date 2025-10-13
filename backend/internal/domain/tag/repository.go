package tag

import "context"

// Repository defines the interface for tag persistence
type Repository interface {
	// Save persists a tag
	Save(ctx context.Context, tag *Tag) error

	// FindByID retrieves a tag by its ID
	FindByID(ctx context.Context, id TagID) (*Tag, error)

	// FindByName retrieves a tag by its name
	FindByName(ctx context.Context, name string) (*Tag, error)

	// FindByCategory retrieves all tags in a category
	FindByCategory(ctx context.Context, category TagCategory) ([]*Tag, error)

	// FindAll retrieves all tags
	FindAll(ctx context.Context) ([]*Tag, error)

	// FindByIDs retrieves multiple tags by their IDs
	FindByIDs(ctx context.Context, ids []TagID) ([]*Tag, error)

	// Delete removes a tag
	Delete(ctx context.Context, id TagID) error

	// Exists checks if a tag with the given name already exists
	Exists(ctx context.Context, name string) (bool, error)
}
