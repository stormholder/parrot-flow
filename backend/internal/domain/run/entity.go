package run

import (
	"errors"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
	"time"
)

type RunID struct {
	shared.ID
}

func NewRunID(value string) (RunID, error) {
	id, err := shared.NewID(value)
	if err != nil {
		return RunID{}, err
	}
	return RunID{ID: id}, nil
}

type Run struct {
	Id         RunID
	ScenarioID scenario.ScenarioID
	Status     shared.Status
	Parameters string
	StartedAt  *shared.Timestamp
	FinishedAt *shared.Timestamp
	CreatedAt  shared.Timestamp
	UpdatedAt  shared.Timestamp
}

func NewRun(id RunID, scenarioID scenario.ScenarioID, parameters string) (*Run, error) {
	if parameters == "" {
		return nil, errors.New("run parameters cannot be empty")
	}

	run := &Run{
		Id:         id,
		ScenarioID: scenarioID,
		Status:     shared.StatusPending,
		Parameters: parameters,
		CreatedAt:  shared.NewTimestamp(time.Now()),
		UpdatedAt:  shared.NewTimestamp(time.Now()),
	}

	return run, nil
}

func (r *Run) Start() error {
	if r.Status != shared.StatusPending {
		return errors.New("can only start a pending run")
	}

	r.Status = shared.StatusRunning
	startedAt := shared.NewTimestamp(time.Now())
	r.StartedAt = &startedAt
	r.UpdatedAt = shared.NewTimestamp(time.Now())

	return nil
}

func (r *Run) Complete() error {
	if r.Status != shared.StatusRunning {
		return errors.New("can only complete a running run")
	}

	r.Status = shared.StatusCompleted
	finishedAt := shared.NewTimestamp(time.Now())
	r.FinishedAt = &finishedAt
	r.UpdatedAt = shared.NewTimestamp(time.Now())

	return nil
}

func (r *Run) Fail(reason string) error {
	if r.Status != shared.StatusRunning {
		return errors.New("can only fail a running run")
	}

	r.Status = shared.StatusFailed
	finishedAt := shared.NewTimestamp(time.Now())
	r.FinishedAt = &finishedAt
	r.UpdatedAt = shared.NewTimestamp(time.Now())

	return nil
}

func (r *Run) Cancel() error {
	if r.Status == shared.StatusCompleted || r.Status == shared.StatusFailed {
		return errors.New("cannot cancel a completed or failed run")
	}

	r.Status = shared.StatusCancelled
	finishedAt := shared.NewTimestamp(time.Now())
	r.FinishedAt = &finishedAt
	r.UpdatedAt = shared.NewTimestamp(time.Now())

	return nil
}
