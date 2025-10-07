package ports

import (
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/models"
	"strconv"
)

func RunParseID(id string) uint64 {
	return parseID(id)
}

func parseID(id string) uint64 {
	if id == "" {
		return 0
	}
	val, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func formatID(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func RunDomainEntityToPersistence(run *run.Run) (*models.ScenarioRun, error) {
	model := &models.ScenarioRun{
		Model: models.Model{
			ID:        parseID(run.Id.String()),
			CreatedAt: run.CreatedAt.Time(),
			UpdatedAt: run.UpdatedAt.Time(),
		},
		ScenarioID: parseID(run.ScenarioID.String()),
		Status:     run.Status.String(),
		Parameters: run.Parameters,
	}

	if run.StartedAt != nil {
		model.StartedAt = run.StartedAt.Time()
	}
	if run.FinishedAt != nil {
		model.FinishedAt = run.FinishedAt.Time()
	}
	return model, nil
}

func RunPersistenceToDomainEntity(model *models.ScenarioRun) (*run.Run, error) {
	runID, err := run.NewRunID(formatID(model.ID))
	if err != nil {
		return nil, err
	}

	scenarioID, err := scenario.NewScenarioID(formatID(model.ScenarioID))
	if err != nil {
		return nil, err
	}

	status, err := shared.NewStatus(model.Status)
	if err != nil {
		return nil, err
	}

	run, err := run.NewRun(runID, scenarioID, model.Parameters)
	if err != nil {
		return nil, err
	}

	run.Status = status
	if !model.StartedAt.IsZero() {
		startedAt := shared.NewTimestamp(model.StartedAt)
		run.StartedAt = &startedAt
	}
	if !model.FinishedAt.IsZero() {
		finishedAt := shared.NewTimestamp(model.FinishedAt)
		run.FinishedAt = &finishedAt
	}

	return run, nil
}
