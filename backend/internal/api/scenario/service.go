package scenario

import (
	"encoding/json"
	"parrotflow/internal/models"
	"parrotflow/pkg/shared"
)

type ScenarioService struct {
	store *ScenarioStore
}

func NewScenarioService(store *ScenarioStore) *ScenarioService {
	return &ScenarioService{store}
}

func (s *ScenarioService) Create() (models.Scenario, error) {
	var blocks []Node = make([]Node, 1)
	newId := shared.CustomUUID()
	blocks[0] = Node{
		ID: newId,
		Position: Point2D{
			X: 0,
			Y: 0,
		},
	}
	var edges []Edge = make([]Edge, 0)
	payload := ScenarioPayload{
		Blocks: blocks,
		Edges:  edges,
	}
	byteStr, err := json.Marshal(payload)

	if err != nil {
		return models.Scenario{}, err
	}

	payloadStr := string(byteStr)

	newScenario := models.Scenario{
		ScenarioBase: models.ScenarioBase{
			Name: "New Scenario",
		},
		Payload:    payloadStr,
		InputData:  "",
		Parameters: "",
	}

	return newScenario, nil
}
