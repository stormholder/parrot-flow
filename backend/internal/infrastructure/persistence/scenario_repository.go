package persistence

import (
	"context"
	"errors"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/models"
	"parrotflow/internal/ports"

	"gorm.io/gorm"
)

type ScenarioRepository struct {
	db *gorm.DB
}

func NewScenarioRepository(db *gorm.DB) *ScenarioRepository {
	return &ScenarioRepository{db: db}
}

func (r *ScenarioRepository) Save(ctx context.Context, s *scenario.Scenario) error {
	model, err := ports.ScenarioDomainEntityToPersistence(s)

	if err = r.db.WithContext(ctx).Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *ScenarioRepository) FindByID(ctx context.Context, id scenario.ScenarioID) (*scenario.Scenario, error) {
	var model models.Scenario
	if err := r.db.WithContext(ctx).Where("id = ?", ports.ScenarioParseID(id.String())).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("scenario not found")
		}
		return nil, err
	}

	return ports.ScenarioPersistenceToDomainEntity(&model)
}

func (r *ScenarioRepository) FindByName(ctx context.Context, name string) (*scenario.Scenario, error) {
	var model models.Scenario
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("scenario not found")
		}
		return nil, err
	}

	return ports.ScenarioPersistenceToDomainEntity(&model)
}

func (r *ScenarioRepository) FindAll(ctx context.Context, criteria scenario.SearchCriteria) ([]*scenario.Scenario, error) {
	var models []models.Scenario
	query := r.db.WithContext(ctx)

	if criteria.Name != "" {
		query = query.Where("name LIKE ?", "%"+criteria.Name+"%")
	}
	if criteria.Tag != "" {
		query = query.Where("tag = ?", criteria.Tag)
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

	scenarios := make([]*scenario.Scenario, len(models))
	for i, model := range models {
		scenario, err := ports.ScenarioPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		scenarios[i] = scenario
	}

	return scenarios, nil
}

func (r *ScenarioRepository) Delete(ctx context.Context, id scenario.ScenarioID) error {
	return r.db.WithContext(ctx).Where("id = ?", ports.ScenarioParseID(id.String())).Delete(&models.Scenario{}).Error
}

func (r *ScenarioRepository) Exists(ctx context.Context, id scenario.ScenarioID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Scenario{}).Where("id = ?", ports.ScenarioParseID(id.String())).Count(&count).Error
	return count > 0, err
}
