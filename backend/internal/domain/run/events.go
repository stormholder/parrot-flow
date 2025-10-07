package run

import (
	"parrotflow/internal/domain/shared"
	"time"
)

var (
	EventRunCreated   = "RunCreated"
	EventRunStarted   = "RunStarted"
	EventRunCompleted = "RunCompleted"
	EventRunFailed    = "RunFailed"
	EventRunCancelled = "RunCancelled"
)

type RunCreated struct {
	shared.BaseEvent
	RunID      string
	ScenarioID string
	Parameters string
}

type RunStarted struct {
	shared.BaseEvent
	RunID      string
	ScenarioID string
	StartedAt  time.Time
}

type RunCompleted struct {
	shared.BaseEvent
	RunID      string
	ScenarioID string
	FinishedAt time.Time
}

type RunFailed struct {
	shared.BaseEvent
	RunID      string
	ScenarioID string
	Reason     string
	FailedAt   time.Time
}

type RunCancelled struct {
	shared.BaseEvent
	RunID       string
	ScenarioID  string
	CancelledAt time.Time
}
