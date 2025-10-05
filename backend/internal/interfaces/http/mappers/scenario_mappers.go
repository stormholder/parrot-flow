package mappers

import (
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/interfaces/http/commands"
	"parrotflow/internal/interfaces/http/queries"
)

func ToGetScenarioResponse(scenario *scenario.Scenario) *queries.GetScenarioResponse {
	response := &queries.GetScenarioResponse{}
	response.Body.ID = scenario.Id.String()
	response.Body.Name = scenario.Name
	response.Body.Description = scenario.Description
	response.Body.Tag = scenario.Tag
	response.Body.Icon = scenario.Icon
	response.Body.CreatedAt = scenario.CreatedAt.Time().Format("2006-01-02T15:04:05Z")
	response.Body.UpdatedAt = scenario.UpdatedAt.Time().Format("2006-01-02T15:04:05Z")

	return response
}

func ToCreateScenarioResponse(scenario *scenario.Scenario) *commands.CreateScenarioResponse {
	response := &commands.CreateScenarioResponse{}
	response.Body.ID = scenario.Id.String()
	response.Body.Name = scenario.Name
	response.Body.Description = scenario.Description
	response.Body.Tag = scenario.Tag
	response.Body.Icon = scenario.Icon
	response.Body.CreatedAt = scenario.CreatedAt.Time().Format("2006-01-02T15:04:05Z")
	response.Body.UpdatedAt = scenario.UpdatedAt.Time().Format("2006-01-02T15:04:05Z")

	return response
}

func toScenarioListItem(scenario *scenario.Scenario) *queries.ScenarioResponseItem {
	item := &queries.ScenarioResponseItem{}

	item.ID = scenario.Id.String()
	item.Name = scenario.Name
	item.Description = scenario.Description
	item.Tag = scenario.Tag
	item.Icon = scenario.Icon
	item.CreatedAt = scenario.CreatedAt.Time().Format("2006-01-02T15:04:05Z")
	item.UpdatedAt = scenario.UpdatedAt.Time().Format("2006-01-02T15:04:05Z")

	return item
}

func ToListScenarioResponse(scenarios []*scenario.Scenario, page int, rpp int) *queries.ListScenariosResponse {

	response := &queries.ListScenariosResponse{}
	response.Body.Total = len(scenarios)
	response.Body.Page = page
	response.Body.RPP = rpp

	response.Body.Data = make([]queries.ScenarioResponseItem, len(scenarios))

	for i, scenario := range scenarios {
		response.Body.Data[i] = *toScenarioListItem(scenario)
	}

	return response
}

func ToScenarioUpdateResponse(scenario *scenario.Scenario) *commands.UpdateScenarioResponse {
	response := &commands.UpdateScenarioResponse{}
	response.Body.ID = scenario.Id.String()
	response.Body.Name = scenario.Name
	response.Body.Description = scenario.Description
	response.Body.Tag = scenario.Tag
	response.Body.Icon = scenario.Icon
	response.Body.UpdatedAt = scenario.UpdatedAt.Time().Format("2006-01-02T15:04:05Z")

	return response
}
