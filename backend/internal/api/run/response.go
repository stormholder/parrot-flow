package run

import (
	"parrotflow/internal/api"
	"parrotflow/internal/api/scenario"
	"parrotflow/internal/models"
	"time"
)

type RunListItemEntity struct {
	models.Model
	ScenarioID uint64    `json:"scenario_id"`
	Status     string    `json:"status"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at,omitempty"`
	Parameters string    `json:"parameters"`
}

type RunEntity struct {
	RunListItemEntity
	Scenario scenario.ScenarioContext `json:"scenario"`
}

type RunListResponse struct {
	Body struct {
		api.Pages
	}
}

type RunResponse struct {
	Body struct {
		RunEntity
	}
}

type RunCommandResponse struct {
	Body struct {
		ID        uint      `json:"id"`
		HandledAt time.Time `json:"handled_at"`
		Success   bool      `json:"success"`
		Detail    string    `json:"detail,omitempty"`
	}
}
