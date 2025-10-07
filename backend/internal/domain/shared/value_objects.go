package shared

import (
	"errors"
	"time"
)

type ID struct {
	value string
}

func NewID(value string) (ID, error) {
	if value == "" {
		return ID{}, errors.New("id cannot be empty")
	}
	return ID{value: value}, nil
}

func (id ID) String() string {
	return id.value
}

func (id ID) IsEmpty() bool {
	return id.value == ""
}

type Timestamp struct {
	value time.Time
}

func NewTimestamp(value time.Time) Timestamp {
	return Timestamp{value: value}
}

func (t Timestamp) Time() time.Time {
	return t.value
}

func (t Timestamp) IsZero() bool {
	return t.value.IsZero()
}

type Status struct {
	value string
}

func NewStatus(value string) (Status, error) {
	if value == "" {
		return Status{}, errors.New("status cannot be empty")
	}
	return Status{value: value}, nil
}

func (s Status) String() string {
	return s.value
}

// Common statuses
var (
	StatusPending   = Status{value: "PENDING"}
	StatusRunning   = Status{value: "RUNNING"}
	StatusCompleted = Status{value: "COMPLETED"}
	StatusFailed    = Status{value: "FAILED"}
	StatusCancelled = Status{value: "CANCELLED"}
)
