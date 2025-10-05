package scenario

import (
	"errors"
	"parrotflow/internal/domain/shared"
	"time"
)

type ScenarioID struct {
	shared.ID
}

func NewScenarioID(value string) (ScenarioID, error) {
	id, err := shared.NewID(value)
	if err != nil {
		return ScenarioID{}, err
	}
	return ScenarioID{ID: id}, nil
}

type Scenario struct {
	Id          ScenarioID
	Name        string
	Description string
	Tag         string
	Icon        string
	Context     Context
	InputData   InputData
	Parameters  Parameters
	CreatedAt   shared.Timestamp
	UpdatedAt   shared.Timestamp
}

func NewScenario(id ScenarioID, name string) (*Scenario, error) {
	if name == "" {
		return nil, errors.New("scenario name cannot be empty")
	}

	scenario := &Scenario{
		Id:        id,
		Name:      name,
		CreatedAt: shared.NewTimestamp(time.Now()),
		UpdatedAt: shared.NewTimestamp(time.Now()),
	}

	return scenario, nil
}

func (s *Scenario) UpdateName(name string) error {
	if name == "" {
		return errors.New("scenario name cannot be empty")
	}
	s.Name = name
	s.UpdatedAt = shared.NewTimestamp(time.Now())
	return nil
}

func (s *Scenario) UpdateDescription(description string) {
	s.Description = description
	s.UpdatedAt = shared.NewTimestamp(time.Now())
}

func (s *Scenario) UpdateTag(tag string) {
	s.Tag = tag
	s.UpdatedAt = shared.NewTimestamp(time.Now())
}

func (s *Scenario) UpdateIcon(icon string) {
	s.Icon = icon
	s.UpdatedAt = shared.NewTimestamp(time.Now())
}

func (s *Scenario) UpdateContext(context Context) {
	s.Context = context
	s.UpdatedAt = shared.NewTimestamp(time.Now())
}

func (s *Scenario) UpdateInputData(inputData InputData) {
	s.InputData = inputData
	s.UpdatedAt = shared.NewTimestamp(time.Now())
}

func (s *Scenario) UpdateParameters(parameters Parameters) {
	s.Parameters = parameters
	s.UpdatedAt = shared.NewTimestamp(time.Now())
}
