package scenario

import (
	"fmt"
	"math"
	"parrotflow/internal/api"
	"parrotflow/internal/models"

	"gorm.io/gorm"
)

type ScenarioStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *ScenarioStore {
	return &ScenarioStore{db}
}

func (s *ScenarioStore) List(query ScenarioQuery) (api.Pages, error) {
	var list []models.Scenario
	fields := []string{"id", "created_at", "updated_at", "name", "description", "tag", "icon"}
	var count int64
	err := s.db.Model(&models.Scenario{}).Count(&count).Error
	if err != nil {
		return api.Pages{}, err
	}

	ctx := api.BuildQuery(s.db.Model(&models.Scenario{}), api.GenericQuery{PageQuery: query.PageQuery, OrderByQuery: query.OrderByQuery}, fields)

	if query.Name != "" {
		ctx = ctx.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if len(query.Tags) > 0 {
		ctx = ctx.Where("tag IN ?", query.Tags)
	}

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

func (s *ScenarioStore) GetByID(id uint) (models.Scenario, error) {
	var scenario models.Scenario
	err := s.db.
		Where("id = ?", id).
		First(&scenario).
		Error
	return scenario, err
}

func (s *ScenarioStore) Create(scenario models.Scenario) (models.Scenario, error) {
	if err := s.db.Create(&scenario).Error; err != nil {
		return scenario, err
	}
	return scenario, nil
}

func (s *ScenarioStore) Update(scenario models.Scenario) (models.Scenario, error) {
	if err := s.db.Save(&scenario).Error; err != nil {
		return scenario, err
	}
	return scenario, nil
}

func (s *ScenarioStore) Delete(id uint) {
	s.db.Delete(&models.Scenario{}, id)
}
