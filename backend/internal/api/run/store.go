package run

import (
	"fmt"
	"math"
	"parrotflow/internal/api"
	"parrotflow/internal/models"

	"gorm.io/gorm"
)

type RunStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *RunStore {
	return &RunStore{db}
}

func (s *RunStore) List(query RunQuery) (api.Pages, error) {
	var list []models.ScenarioRun
	fields := []string{"id", "created_at", "updated_at", "finished_at", "successful"}
	var count int64
	err := s.db.Model(&models.ScenarioRun{}).Count(&count).Error
	if err != nil {
		return api.Pages{}, err
	}

	ctx := api.BuildQuery(s.db.Model(&models.ScenarioRun{}), api.GenericQuery{PageQuery: query.PageQuery, OrderByQuery: query.OrderByQuery}, fields)

	var pages int
	if count <= int64(query.PerPage) {
		pages = 1
	} else {
		pages = int(math.Floor(float64(count) / float64(query.PerPage)))
	}

	err = ctx.Find(&list).Error
	if err != nil {
		return api.Pages{}, err
	}
	fmt.Println("Scenarios found: ", len(list))
	return api.Pages{
		TotalItems:  int(count),
		CurrentPage: query.Page,
		RowsPerPage: query.PerPage,
		TotalPages:  pages,
		Data:        list,
	}, nil
}

func (s *RunStore) GetByID(id uint) (models.ScenarioRun, error) {
	var run models.ScenarioRun
	err := s.db.
		Where("id = ?", id).
		First(&run).
		Error
	return run, err
}

func (s *RunStore) Create(run models.ScenarioRun) (models.ScenarioRun, error) {
	if err := s.db.Create(&run).Error; err != nil {
		return run, err
	}
	return run, nil
}

func (s *RunStore) Update(run models.ScenarioRun) (models.ScenarioRun, error) {
	if err := s.db.Save(&run).Error; err != nil {
		return run, err
	}
	return run, nil
}
