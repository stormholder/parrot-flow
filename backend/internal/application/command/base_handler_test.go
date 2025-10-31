package command

import (
	"errors"
	"parrotflow/internal/domain/shared"
	"testing"
)

// Mock EventBus for testing
type MockEventBus struct {
	publishedEvents []shared.DomainEvent
	publishError    error
}

func (m *MockEventBus) Publish(event shared.DomainEvent) error {
	if m.publishError != nil {
		return m.publishError
	}
	m.publishedEvents = append(m.publishedEvents, event)
	return nil
}

func (m *MockEventBus) Subscribe(handler shared.EventHandler) error {
	return nil // Not used in these tests
}

// Mock EventCarrier for testing
type MockEntity struct {
	Events      []shared.DomainEvent
	eventCleared bool
}

func (m *MockEntity) ClearEvents() {
	m.Events = nil
	m.eventCleared = true
}

// Mock DomainEvent for testing
type MockDomainEvent struct {
	shared.BaseEvent
	Data string
}

func TestPublishDomainEvents(t *testing.T) {
	// Test case: Successfully publish multiple events
	mockBus := &MockEventBus{}
	entity := &MockEntity{
		Events: []shared.DomainEvent{
			&MockDomainEvent{Data: "event1"},
			&MockDomainEvent{Data: "event2"},
			&MockDomainEvent{Data: "event3"},
		},
	}

	PublishDomainEvents(mockBus, entity.Events, entity)

	// Verify all events were published
	if len(mockBus.publishedEvents) != 3 {
		t.Errorf("Expected 3 events published, got %d", len(mockBus.publishedEvents))
	}

	// Verify events were cleared
	if !entity.eventCleared {
		t.Error("Events were not cleared from entity")
	}
}

func TestPublishDomainEvents_EmptyEvents(t *testing.T) {
	// Test case: No events to publish
	mockBus := &MockEventBus{}
	entity := &MockEntity{
		Events: []shared.DomainEvent{},
	}

	PublishDomainEvents(mockBus, entity.Events, entity)

	// Verify no events published
	if len(mockBus.publishedEvents) != 0 {
		t.Errorf("Expected 0 events published, got %d", len(mockBus.publishedEvents))
	}

	// Verify ClearEvents was still called
	if !entity.eventCleared {
		t.Error("ClearEvents should be called even with no events")
	}
}

func TestPublishDomainEvents_PublishErrorHandled(t *testing.T) {
	// Test case: Error during publishing should not panic (errors are ignored)
	mockBus := &MockEventBus{
		publishError: errors.New("test error"),
	}
	entity := &MockEntity{
		Events: []shared.DomainEvent{
			&MockDomainEvent{Data: "event1"},
		},
	}

	// Should not panic
	PublishDomainEvents(mockBus, entity.Events, entity)

	// Events should still be cleared despite error
	if !entity.eventCleared {
		t.Error("Events should be cleared even when publish fails")
	}
}
