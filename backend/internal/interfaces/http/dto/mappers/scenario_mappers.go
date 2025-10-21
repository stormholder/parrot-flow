package mappers

import (
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

func buildScenarioDTO(s *scenario.Scenario) queries.ScenarioResponseItem {
	return queries.ScenarioResponseItem{
		ID:          s.Id.String(),
		Name:        s.Name,
		Description: s.Description,
		Tag:         s.Tag,
		Icon:        s.Icon,
		CreatedAt:   FormatTimestamp(s.CreatedAt.Time()),
		UpdatedAt:   FormatTimestamp(s.UpdatedAt.Time()),
	}
}

type ScenarioCreateMapper struct{}

func (m ScenarioCreateMapper) Map(s *scenario.Scenario) *commands.CreateScenarioResponse {
	dto := buildScenarioDTO(s)
	response := &commands.CreateScenarioResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Description = dto.Description
	response.Body.Tag = dto.Tag
	response.Body.Icon = dto.Icon
	response.Body.CreatedAt = dto.CreatedAt
	response.Body.UpdatedAt = dto.UpdatedAt
	return response
}

type ScenarioUpdateMapper struct{}

func (m ScenarioUpdateMapper) Map(s *scenario.Scenario) *commands.UpdateScenarioResponse {
	dto := buildScenarioDTO(s)
	response := &commands.UpdateScenarioResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Description = dto.Description
	response.Body.Tag = dto.Tag
	response.Body.Icon = dto.Icon
	response.Body.UpdatedAt = dto.UpdatedAt
	return response
}

type ScenarioDeleteMapper struct{}

func (m ScenarioDeleteMapper) Map() *commands.DeleteScenarioResponse {
	response := &commands.DeleteScenarioResponse{}
	response.Body.Success = true
	return response
}

type ScenarioGetMapper struct{}

func (m ScenarioGetMapper) Map(s *scenario.Scenario) *queries.GetScenarioResponse {
	response := &queries.GetScenarioResponse{}
	response.Body = buildScenarioDTO(s)
	return response
}

type ScenarioListMapper struct {
	Page int
	RPP  int
}

func (m ScenarioListMapper) Map(scenarios []*scenario.Scenario) *queries.ListScenariosResponse {
	response := &queries.ListScenariosResponse{}
	response.Body.Total = len(scenarios)
	response.Body.Page = m.Page
	response.Body.RPP = m.RPP
	response.Body.Data = MapSlicePtr(scenarios, buildScenarioDTO)
	return response
}

func ToCreateScenarioResponse(s *scenario.Scenario) *commands.CreateScenarioResponse {
	return ScenarioCreateMapper{}.Map(s)
}

func ToScenarioUpdateResponse(s *scenario.Scenario) *commands.UpdateScenarioResponse {
	return ScenarioUpdateMapper{}.Map(s)
}

func ToGetScenarioResponse(s *scenario.Scenario) *queries.GetScenarioResponse {
	return ScenarioGetMapper{}.Map(s)
}

func ToListScenarioResponse(scenarios []*scenario.Scenario, page int, rpp int) *queries.ListScenariosResponse {
	return ScenarioListMapper{Page: page, RPP: rpp}.Map(scenarios)
}
