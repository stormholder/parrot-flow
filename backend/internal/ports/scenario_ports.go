package ports

import (
	"encoding/json"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/models"
)

func ScenarioParseID(id string) uint64 {
	return parseID(id)
}

func ScenarioDomainEntityToPersistence(s *scenario.Scenario) (*models.Scenario, error) {
	model := &models.Scenario{
		ScenarioBase: models.ScenarioBase{
			Model: models.Model{
				ID:        parseID(s.Id.String()),
				CreatedAt: s.CreatedAt.Time(),
				UpdatedAt: s.UpdatedAt.Time(),
			},
			Name:        s.Name,
			Description: s.Description,
			Tag:         s.Tag,
			Icon:        s.Icon,
		},
		Context:    marshalContext(s.Context),
		InputData:  marshalInputData(s.InputData),
		Parameters: marshalParameters(s.Parameters),
	}
	return model, nil
}

func ScenarioPersistenceToDomainEntity(model *models.Scenario) (*scenario.Scenario, error) {
	scenarioID, err := scenario.NewScenarioID(formatID(model.ID))
	if err != nil {
		return nil, err
	}

	s, err := scenario.NewScenario(scenarioID, model.Name)
	if err != nil {
		return nil, err
	}

	s.UpdateDescription(model.Description)
	s.UpdateTag(model.Tag)
	s.UpdateIcon(model.Icon)

	if model.Context != "" {
		context, err := unmarshalContext(model.Context)
		if err != nil {
			return nil, err
		}
		s.UpdateContext(context)
	}

	if model.InputData != "" {
		inputData, err := unmarshalInputData(model.InputData)
		if err != nil {
			return nil, err
		}
		s.UpdateInputData(inputData)
	}

	if model.Parameters != "" {
		parameters, err := unmarshalParameters(model.Parameters)
		if err != nil {
			return nil, err
		}
		s.UpdateParameters(parameters)
	}

	return s, nil
}

func marshalContext(context scenario.Context) string {
	data, _ := json.Marshal(context)
	return string(data)
}

func marshalInputData(inputData scenario.InputData) string {
	data, _ := json.Marshal(inputData)
	return string(data)
}

func marshalParameters(parameters scenario.Parameters) string {
	data, _ := json.Marshal(parameters)
	return string(data)
}

func unmarshalContext(data string) (scenario.Context, error) {
	var context scenario.Context
	err := json.Unmarshal([]byte(data), &context)
	return context, err
}

func unmarshalInputData(data string) (scenario.InputData, error) {
	var inputData scenario.InputData
	err := json.Unmarshal([]byte(data), &inputData)
	return inputData, err
}

func unmarshalParameters(data string) (scenario.Parameters, error) {
	var parameters scenario.Parameters
	err := json.Unmarshal([]byte(data), &parameters)
	return parameters, err
}
