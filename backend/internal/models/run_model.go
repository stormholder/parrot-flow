package models

import "time"

type ScenarioRun struct {
	Model
	ScenarioID uint64    `json:"scenario_id" gorm:"not null"`
	Status     string    `json:"status" gorm:"not null"`
	StartedAt  time.Time `json:"started_at" gorm:"not null"`
	FinishedAt time.Time `json:"finished_at,omitempty"`
	Parameters string    `json:"parameters" gorm:"not null"`
}
