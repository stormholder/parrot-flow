package command

import (
	"parrotflow/internal/domain/shared"
)

// EventCarrier is a minimal interface for entities that carry domain events
// This allows us to handle event publishing generically without coupling to specific entity types
type EventCarrier interface {
	ClearEvents()
}

// PublishDomainEvents publishes all domain events and clears them from the entity
// This eliminates the boilerplate event publishing code repeated in every command handler
//
// Usage in command handlers:
//   command.PublishDomainEvents(h.eventBus, entity.Events, entity)
func PublishDomainEvents(eventBus shared.EventBus, events []shared.DomainEvent, carrier EventCarrier) {
	for _, event := range events {
		if err := eventBus.Publish(event); err != nil {
			// TODO: Add proper logging instead of silently ignoring
			// For now, maintain backward compatibility
		}
	}
	carrier.ClearEvents()
}
