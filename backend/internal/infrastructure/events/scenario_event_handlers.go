package events

import (
	"log"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
)

// ScenarioCreatedHandler handles scenario created events
type ScenarioCreatedHandler struct{}

// NewScenarioCreatedHandler creates a new scenario created handler
func NewScenarioCreatedHandler() *ScenarioCreatedHandler {
	return &ScenarioCreatedHandler{}
}

// Handle handles the scenario created event
func (h *ScenarioCreatedHandler) Handle(event shared.DomainEvent) error {
	if scenarioCreated, ok := event.(scenario.ScenarioCreated); ok {
		log.Printf("Scenario created: %s with name: %s", scenarioCreated.ScenarioID, scenarioCreated.Name)
		// Here you could add additional logic like:
		// - Send notifications
		// - Update search indexes
		// - Trigger other processes
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *ScenarioCreatedHandler) CanHandle(eventType string) bool {
	return eventType == "ScenarioCreated"
}

// ScenarioUpdatedHandler handles scenario updated events
type ScenarioUpdatedHandler struct{}

// NewScenarioUpdatedHandler creates a new scenario updated handler
func NewScenarioUpdatedHandler() *ScenarioUpdatedHandler {
	return &ScenarioUpdatedHandler{}
}

// Handle handles the scenario updated event
func (h *ScenarioUpdatedHandler) Handle(event shared.DomainEvent) error {
	if scenarioUpdated, ok := event.(scenario.ScenarioUpdated); ok {
		log.Printf("Scenario updated: %s with changes: %v", scenarioUpdated.ScenarioID, scenarioUpdated.Changes)
		// Here you could add additional logic like:
		// - Update search indexes
		// - Send notifications
		// - Trigger cache invalidation
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *ScenarioUpdatedHandler) CanHandle(eventType string) bool {
	return eventType == "ScenarioUpdated"
}

// ScenarioDeletedHandler handles scenario deleted events
type ScenarioDeletedHandler struct{}

// NewScenarioDeletedHandler creates a new scenario deleted handler
func NewScenarioDeletedHandler() *ScenarioDeletedHandler {
	return &ScenarioDeletedHandler{}
}

// Handle handles the scenario deleted event
func (h *ScenarioDeletedHandler) Handle(event shared.DomainEvent) error {
	if scenarioDeleted, ok := event.(scenario.ScenarioDeleted); ok {
		log.Printf("Scenario deleted: %s", scenarioDeleted.ScenarioID)
		// Here you could add additional logic like:
		// - Clean up related data
		// - Update search indexes
		// - Send notifications
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *ScenarioDeletedHandler) CanHandle(eventType string) bool {
	return eventType == "ScenarioDeleted"
}
