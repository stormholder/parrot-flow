package scenario

import (
	"encoding/json"
	"parrotflow/internal/api"
	"parrotflow/internal/models"
	"parrotflow/pkg/shared"
)

type ScenarioService struct {
	store *ScenarioStore
}

func NewScenarioService(store *ScenarioStore) *ScenarioService {
	return &ScenarioService{store}
}

func (s *ScenarioService) FindMany(query ScenarioQuery) (api.Pages, error) {
	return s.store.List(query)
}

func (s *ScenarioService) FindOne(id uint) (models.Scenario, error) {
	return s.store.GetByID(id)
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
	context := ScenarioContext{
		Blocks: blocks,
		Edges:  edges,
	}
	byteStr, err := json.Marshal(context)

	if err != nil {
		return models.Scenario{}, err
	}

	payloadStr := string(byteStr)

	newScenario := models.Scenario{
		ScenarioBase: models.ScenarioBase{
			Name: "New Scenario",
		},
		Context:    payloadStr,
		InputData:  "",
		Parameters: "",
	}

	return s.store.Create(newScenario)
}

func (s *ScenarioService) Update(id uint, patch ScenarioPatch) (*models.Scenario, error) {
	scenario, err := s.store.GetByID(id)
	if err != nil {
		return nil, err
	}

	if patch.Name != nil {
		scenario.Name = *patch.Name
	}
	if patch.Description != nil {
		scenario.Description = *patch.Description
	}
	if patch.Icon != nil {
		scenario.Icon = *patch.Icon
	}
	if patch.Tag != nil {
		scenario.Tag = *patch.Tag
	}
	if patch.Context != nil {
		jsonString, err := json.Marshal(patch.Context)
		if err != nil {
			return nil, err
		}
		scenario.Context = string(jsonString)
	}
	if patch.InputData != nil {
		jsonString, err := json.Marshal(patch.InputData)
		if err != nil {
			return nil, err
		}
		scenario.InputData = string(jsonString)
	}
	if patch.Parameters != nil {
		jsonString, err := json.Marshal(patch.Parameters)
		if err != nil {
			return nil, err
		}
		scenario.Parameters = string(jsonString)
	}

	updatedScenario, err := s.store.Update(scenario)
	if err != nil {
		return nil, err
	}
	return &updatedScenario, nil
}

func (s *ScenarioService) Delete(id uint) (*struct{}, error) {
	if err := s.store.Delete(id); err != nil {
		return nil, err
	}
	return nil, nil
}
