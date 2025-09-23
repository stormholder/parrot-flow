package run

import (
	"parrotflow/internal/api"
	"parrotflow/internal/models"
)

type RunService struct {
	store *RunStore
}

func NewRunService(store *RunStore) *RunService {
	return &RunService{store}
}

func (s *RunService) FindMany(query RunQuery) (api.Pages, error) {
	return s.store.List(query)
}

func (s *RunService) FindOne(id uint) (models.ScenarioRun, error) {
	return s.store.GetByID(id)
}
