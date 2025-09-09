package scenario

import "parrotflow/internal/models"

type ScenarioPayload struct {
	Blocks []Node `json:"blocks"`
	Edges  []Edge `json:"edges"`
}

type ScenarioResponse struct {
	models.ScenarioBase
	Payload ScenarioPayload `json:"payload"`
}
