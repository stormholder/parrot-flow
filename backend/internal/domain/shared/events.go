package shared

import (
	"parrotflow/pkg/shared"
	"time"
)

type DomainEvent interface {
	EventID() string
	EventType() string
	OccurredAt() time.Time
	AggregateID() string
}
type BaseEvent struct {
	eventID     string
	eventType   string
	occurredAt  time.Time
	aggregateID string
}

func NewBaseEvent(eventType, aggregateID string) BaseEvent {
	return BaseEvent{
		eventID:     generateEventID(),
		eventType:   eventType,
		occurredAt:  time.Now(),
		aggregateID: aggregateID,
	}
}

func (e BaseEvent) EventID() string {
	return e.eventID
}

func (e BaseEvent) EventType() string {
	return e.eventType
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.occurredAt
}

func (e BaseEvent) AggregateID() string {
	return e.aggregateID
}

type EventHandler interface {
	Handle(event DomainEvent) error
	CanHandle(eventType string) bool
}
type EventBus interface {
	Publish(event DomainEvent) error
	Subscribe(handler EventHandler) error
}

func generateEventID() string {
	return shared.CustomUUID()
}
