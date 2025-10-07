package events

import (
	"log"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/shared"
)

// RunCreatedHandler handles run created events
type RunCreatedHandler struct{}

// NewRunCreatedHandler creates a new run created handler
func NewRunCreatedHandler() *RunCreatedHandler {
	return &RunCreatedHandler{}
}

// Handle handles the run created event
func (h *RunCreatedHandler) Handle(event shared.DomainEvent) error {
	if runCreated, ok := event.(run.RunCreated); ok {
		log.Printf("Run created: %s for scenario: %s", runCreated.RunID, runCreated.ScenarioID)
		// Here you could add additional logic like:
		// - Send notifications
		// - Update metrics
		// - Trigger agent execution
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *RunCreatedHandler) CanHandle(eventType string) bool {
	return eventType == "RunCreated"
}

// RunStartedHandler handles run started events
type RunStartedHandler struct{}

// NewRunStartedHandler creates a new run started handler
func NewRunStartedHandler() *RunStartedHandler {
	return &RunStartedHandler{}
}

// Handle handles the run started event
func (h *RunStartedHandler) Handle(event shared.DomainEvent) error {
	if runStarted, ok := event.(run.RunStarted); ok {
		log.Printf("Run started: %s for scenario: %s at %v", runStarted.RunID, runStarted.ScenarioID, runStarted.StartedAt)
		// Here you could add additional logic like:
		// - Send notifications
		// - Update metrics
		// - Trigger monitoring
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *RunStartedHandler) CanHandle(eventType string) bool {
	return eventType == "RunStarted"
}

// RunCompletedHandler handles run completed events
type RunCompletedHandler struct{}

// NewRunCompletedHandler creates a new run completed handler
func NewRunCompletedHandler() *RunCompletedHandler {
	return &RunCompletedHandler{}
}

// Handle handles the run completed event
func (h *RunCompletedHandler) Handle(event shared.DomainEvent) error {
	if runCompleted, ok := event.(run.RunCompleted); ok {
		log.Printf("Run completed: %s for scenario: %s at %v", runCompleted.RunID, runCompleted.ScenarioID, runCompleted.FinishedAt)
		// Here you could add additional logic like:
		// - Send notifications
		// - Update metrics
		// - Trigger cleanup
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *RunCompletedHandler) CanHandle(eventType string) bool {
	return eventType == "RunCompleted"
}

// RunFailedHandler handles run failed events
type RunFailedHandler struct{}

// NewRunFailedHandler creates a new run failed handler
func NewRunFailedHandler() *RunFailedHandler {
	return &RunFailedHandler{}
}

// Handle handles the run failed event
func (h *RunFailedHandler) Handle(event shared.DomainEvent) error {
	if runFailed, ok := event.(run.RunFailed); ok {
		log.Printf("Run failed: %s for scenario: %s at %v with reason: %s", runFailed.RunID, runFailed.ScenarioID, runFailed.FailedAt, runFailed.Reason)
		// Here you could add additional logic like:
		// - Send notifications
		// - Update metrics
		// - Trigger retry logic
		// - Send alerts
	}
	return nil
}

// CanHandle checks if this handler can handle the event type
func (h *RunFailedHandler) CanHandle(eventType string) bool {
	return eventType == "RunFailed"
}
