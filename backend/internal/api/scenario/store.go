package scenario

import (
	"gorm.io/gorm"
)

type ScenarioStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *ScenarioStore {
	return &ScenarioStore{db}
}
