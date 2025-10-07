package scenario

import (
	"parrotflow/internal/domain/shared"
)

var (
	EventScenarioCreated           = "ScenarioCreated"
	EventScenarioUpdated           = "ScenarioUpdated"
	EventScenarioDeleted           = "ScenarioDeleted"
	EventScenarioContextUpdated    = "ScenarioContextUpdated"
	EventScenarioParametersUpdated = "ScenarioParametersUpdated"
)

type ScenarioCreated struct {
	shared.BaseEvent
	ScenarioID string
	Name       string
}

type ScenarioUpdated struct {
	shared.BaseEvent
	ScenarioID string
	Changes    map[string]interface{}
}

type ScenarioDeleted struct {
	shared.BaseEvent
	ScenarioID string
}

type ScenarioContextUpdated struct {
	shared.BaseEvent
	ScenarioID string
	Context    Context
}

type ScenarioParametersUpdated struct {
	shared.BaseEvent
	ScenarioID string
	Parameters Parameters
}
