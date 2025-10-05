package persistence

import (
	"context"
	"errors"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/models"
	"parrotflow/internal/ports"

	"gorm.io/gorm"
)

type RunRepository struct {
	db *gorm.DB
}

func NewRunRepository(db *gorm.DB) *RunRepository {
	return &RunRepository{db: db}
}

func (r *RunRepository) Save(ctx context.Context, run *run.Run) error {
	model, err := ports.RunDomainEntityToPersistence(run)

	if err = r.db.WithContext(ctx).Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *RunRepository) FindByID(ctx context.Context, id run.RunID) (*run.Run, error) {
	var model models.ScenarioRun
	if err := r.db.WithContext(ctx).Where("id = ?", ports.RunParseID(id.String())).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("run not found")
		}
		return nil, err
	}

	return ports.RunPersistenceToDomainEntity(&model)
}

func (r *RunRepository) FindByScenarioID(ctx context.Context, scenarioID scenario.ScenarioID) ([]*run.Run, error) {
	var models []models.ScenarioRun
	if err := r.db.WithContext(ctx).Where("scenario_id = ?", ports.RunParseID(scenarioID.String())).Find(&models).Error; err != nil {
		return nil, err
	}

	runs := make([]*run.Run, len(models))
	for i, model := range models {
		run, err := ports.RunPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		runs[i] = run
	}

	return runs, nil
}

func (r *RunRepository) FindAll(ctx context.Context, criteria run.SearchCriteria) ([]*run.Run, error) {
	var models []models.ScenarioRun
	query := r.db.WithContext(ctx)

	if !criteria.ScenarioID.IsEmpty() {
		query = query.Where("scenario_id = ?", ports.RunParseID(criteria.ScenarioID.String()))
	}
	if criteria.Status != "" {
		query = query.Where("status = ?", criteria.Status)
	}

	orderBy := criteria.OrderBy
	if orderBy == "" {
		orderBy = "created_at"
	}
	orderDir := criteria.OrderDir
	if orderDir == "" {
		orderDir = "desc"
	}
	query = query.Order(orderBy + " " + orderDir)

	if criteria.Limit > 0 {
		query = query.Limit(criteria.Limit)
	}
	if criteria.Offset > 0 {
		query = query.Offset(criteria.Offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	runs := make([]*run.Run, len(models))
	for i, model := range models {
		run, err := ports.RunPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		runs[i] = run
	}

	return runs, nil
}

func (r *RunRepository) Delete(ctx context.Context, id run.RunID) error {
	return r.db.WithContext(ctx).Where("id = ?", ports.RunParseID(id.String())).Delete(&models.ScenarioRun{}).Error
}

func (r *RunRepository) Exists(ctx context.Context, id run.RunID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ScenarioRun{}).Where("id = ?", ports.RunParseID(id.String())).Count(&count).Error
	return count > 0, err
}
