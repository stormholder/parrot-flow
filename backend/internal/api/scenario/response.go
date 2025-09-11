package scenario

import (
	"parrotflow/internal/api"
	"parrotflow/internal/models"
)

type ScenarioResponse struct {
	Body struct {
		models.ScenarioBase
		Payload ScenarioContext `json:"context"`
	}
}

type ScenarioListResponse struct {
	Body struct {
		api.Pages
	}
}

type ScenarioCreateResponse struct {
	Body struct {
		models.Scenario
	}
}
