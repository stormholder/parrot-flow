package tag

import "parrotflow/internal/domain/shared"

const (
	EventTagCreated = "tag.created"
	EventTagDeleted = "tag.deleted"
)

type TagCreated struct {
	shared.BaseEvent
	TagID       string
	Name        string
	Category    string
	Description string
}

type TagDeleted struct {
	shared.BaseEvent
	TagID string
	Name  string
}
