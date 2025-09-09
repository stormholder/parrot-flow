package scenario

import "parrotflow/internal/models"

type ScenarioResponse struct {
	models.ScenarioBase
	Payload ScenarioPayload `json:"payload"`
}
