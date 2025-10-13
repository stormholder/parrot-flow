package tag

import (
	"errors"
	"parrotflow/internal/domain/shared"
	"strings"
	"time"
)

// Domain errors
var (
	ErrTagAlreadyExists = errors.New("tag with this name already exists")
	ErrTagNotFound      = errors.New("tag not found")
)

type TagID struct {
	shared.ID
}

func NewTagID(value string) (TagID, error) {
	id, err := shared.NewID(value)
	if err != nil {
		return TagID{}, err
	}
	return TagID{ID: id}, nil
}

// TagCategory represents the category/type of the tag
type TagCategory struct {
	value string
}

func NewTagCategory(value string) (TagCategory, error) {
	if value == "" {
		return TagCategory{}, errors.New("tag category cannot be empty")
	}
	return TagCategory{value: strings.ToLower(value)}, nil
}

func (tc TagCategory) String() string {
	return tc.value
}

// Common tag categories
var (
	CategorySystem   = TagCategory{value: "system"}   // Auto-generated tags (os, browser, etc.)
	CategoryCustom   = TagCategory{value: "custom"}   // User-defined tags
	CategoryCapacity = TagCategory{value: "capacity"} // Performance-related tags (fast, slow, etc.)
	CategoryRegion   = TagCategory{value: "region"}   // Geographic tags
)

// Tag represents a label that can be attached to various entities
type Tag struct {
	Id          TagID
	Name        string
	Category    TagCategory
	Description string
	Color       string // Hex color for UI display (e.g., "#FF5733")
	IsSystem    bool   // System tags cannot be deleted by users
	CreatedAt   shared.Timestamp
	UpdatedAt   shared.Timestamp
	Events      []shared.DomainEvent
}

// NewTag creates a new tag
func NewTag(id TagID, name string, category TagCategory) (*Tag, error) {
	if name == "" {
		return nil, errors.New("tag name cannot be empty")
	}

	// Normalize tag name (lowercase, trim spaces)
	normalizedName := strings.ToLower(strings.TrimSpace(name))
	if normalizedName == "" {
		return nil, errors.New("tag name cannot be empty after normalization")
	}

	tag := &Tag{
		Id:        id,
		Name:      normalizedName,
		Category:  category,
		IsSystem:  false,
		CreatedAt: shared.NewTimestamp(time.Now()),
		UpdatedAt: shared.NewTimestamp(time.Now()),
		Events:    make([]shared.DomainEvent, 0),
	}

	tag.addEvent(TagCreated{
		BaseEvent:   shared.NewBaseEvent(EventTagCreated, id.String()),
		TagID:       id.String(),
		Name:        normalizedName,
		Category:    category.String(),
		Description: "",
	})

	return tag, nil
}

// NewSystemTag creates a system tag (cannot be deleted)
func NewSystemTag(id TagID, name string, category TagCategory, description string) (*Tag, error) {
	tag, err := NewTag(id, name, category)
	if err != nil {
		return nil, err
	}

	tag.IsSystem = true
	tag.Description = description

	return tag, nil
}

// UpdateDescription updates the tag description
func (t *Tag) UpdateDescription(description string) {
	t.Description = description
	t.UpdatedAt = shared.NewTimestamp(time.Now())
}

// UpdateColor updates the tag display color
func (t *Tag) UpdateColor(color string) error {
	// Basic hex color validation
	if color != "" && !strings.HasPrefix(color, "#") {
		return errors.New("color must be a valid hex color (e.g., #FF5733)")
	}
	if len(color) != 0 && len(color) != 7 {
		return errors.New("color must be in format #RRGGBB")
	}

	t.Color = color
	t.UpdatedAt = shared.NewTimestamp(time.Now())
	return nil
}

// CanDelete checks if the tag can be deleted
func (t *Tag) CanDelete() bool {
	return !t.IsSystem
}

// Delete marks the tag for deletion (only if it's not a system tag)
func (t *Tag) Delete() error {
	if t.IsSystem {
		return errors.New("cannot delete system tag")
	}

	t.addEvent(TagDeleted{
		BaseEvent: shared.NewBaseEvent(EventTagDeleted, t.Id.String()),
		TagID:     t.Id.String(),
		Name:      t.Name,
	})

	return nil
}

func (t *Tag) addEvent(event shared.DomainEvent) {
	t.Events = append(t.Events, event)
}

func (t *Tag) ClearEvents() {
	t.Events = make([]shared.DomainEvent, 0)
}
