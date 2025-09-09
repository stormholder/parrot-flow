package scenario

import "parrotflow/internal/models"

type ScenarioResponse struct {
	Body struct {
		models.ScenarioBase
		Payload ScenarioPayload `json:"payload"`
	}
}

type ScenarioCreateResponse struct {
	Body struct {
		models.Scenario
	}
}
