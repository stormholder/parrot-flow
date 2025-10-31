package mappers

import (
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
	"parrotflow/internal/interfaces/http/dto/shared"
)

// Domain -> DTO conversion helpers
func mapContextToDTO(ctx scenario.Context) shared.ContextDTO {
	return shared.ContextDTO{
		Blocks: MapSlice(ctx.Blocks, mapNodeToDTO),
		Edges:  MapSlice(ctx.Edges, mapEdgeToDTO),
	}
}

func mapNodeToDTO(n scenario.Node) shared.NodeDTO {
	return shared.NodeDTO{
		ID:       n.Id,
		NodeType: n.NodeType,
		Position: shared.Point2D{X: n.Position.X, Y: n.Position.Y},
	}
}

func mapEdgeToDTO(e scenario.Edge) shared.EdgeDTO {
	return shared.EdgeDTO{
		ID:           e.Id,
		Source:       e.Source,
		Target:       e.Target,
		SourceHandle: e.SourceHandle,
		TargetHandle: e.TargetHandle,
		Condition:    e.Condition,
	}
}

func mapInputDataToDTO(input scenario.InputData) shared.InputDataDTO {
	return shared.InputDataDTO{
		Parameters: MapSlice(input.Parameters, mapNodeParametersToDTO),
	}
}

func mapNodeParametersToDTO(np scenario.NodeParameters) shared.NodeParametersDTO {
	return shared.NodeParametersDTO{
		BlockID: np.BlockID,
		Input:   MapSlice(np.Input, mapParameterToDTO),
		Output:  MapSlice(np.Output, mapParameterToDTO),
	}
}

func mapParameterToDTO(p scenario.Parameter) shared.ParameterDTO {
	return shared.ParameterDTO{
		Name:  p.Name,
		Value: p.Value,
	}
}

func mapParametersToDTO(params scenario.Parameters) shared.ParametersDTO {
	return shared.ParametersDTO{
		Input:  MapSlice(params.Input, mapParameterItemToDTO),
		Output: MapSlice(params.Output, mapParameterItemToDTO),
	}
}

func mapParameterItemToDTO(pi scenario.ParameterItem) shared.ParameterItemDTO {
	return shared.ParameterItemDTO{
		Parameter: mapParameterToDTO(pi.Parameter),
		ParamType: pi.ParamType,
		Values:    pi.Values,
	}
}

func buildScenarioDTO(s *scenario.Scenario) queries.ScenarioResponseItem {
	return queries.ScenarioResponseItem{
		ID:          s.Id.String(),
		Name:        s.Name,
		Description: s.Description,
		Tag:         s.Tag,
		Icon:        s.Icon,
		Context:     mapContextToDTO(s.Context),
		InputData:   mapInputDataToDTO(s.InputData),
		Parameters:  mapParametersToDTO(s.Parameters),
		CreatedAt:   FormatTimestamp(s.CreatedAt.Time()),
		UpdatedAt:   FormatTimestamp(s.UpdatedAt.Time()),
	}
}

// DTO -> Domain conversion helpers
func MapContextFromDTO(dto shared.ContextDTO) scenario.Context {
	return scenario.NewContext(
		MapSlice(dto.Blocks, mapNodeFromDTO),
		MapSlice(dto.Edges, mapEdgeFromDTO),
	)
}

func mapNodeFromDTO(dto shared.NodeDTO) scenario.Node {
	return scenario.Node{
		Id:       dto.ID,
		NodeType: dto.NodeType,
		Position: scenario.Point2D{X: dto.Position.X, Y: dto.Position.Y},
	}
}

func mapEdgeFromDTO(dto shared.EdgeDTO) scenario.Edge {
	return scenario.Edge{
		Id:           dto.ID,
		Source:       dto.Source,
		Target:       dto.Target,
		SourceHandle: dto.SourceHandle,
		TargetHandle: dto.TargetHandle,
		Condition:    dto.Condition,
	}
}

func MapInputDataFromDTO(dto shared.InputDataDTO) scenario.InputData {
	return scenario.NewInputData(MapSlice(dto.Parameters, mapNodeParametersFromDTO))
}

func mapNodeParametersFromDTO(dto shared.NodeParametersDTO) scenario.NodeParameters {
	return scenario.NodeParameters{
		BlockID: dto.BlockID,
		Input:   MapSlice(dto.Input, mapParameterFromDTO),
		Output:  MapSlice(dto.Output, mapParameterFromDTO),
	}
}

func mapParameterFromDTO(dto shared.ParameterDTO) scenario.Parameter {
	return scenario.Parameter{
		Name:  dto.Name,
		Value: dto.Value,
	}
}

func MapParametersFromDTO(dto shared.ParametersDTO) scenario.Parameters {
	return scenario.NewParameters(
		MapSlice(dto.Input, mapParameterItemFromDTO),
		MapSlice(dto.Output, mapParameterItemFromDTO),
	)
}

func mapParameterItemFromDTO(dto shared.ParameterItemDTO) scenario.ParameterItem {
	return scenario.NewParameterItem(
		mapParameterFromDTO(dto.Parameter),
		dto.ParamType,
		dto.Values,
	)
}

// Mapper functions using functional approach

func ScenarioToCreateResponse(s *scenario.Scenario) *commands.CreateScenarioResponse {
	dto := buildScenarioDTO(s)
	response := &commands.CreateScenarioResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Description = dto.Description
	response.Body.Tag = dto.Tag
	response.Body.Icon = dto.Icon
	response.Body.Context = dto.Context
	response.Body.InputData = dto.InputData
	response.Body.Parameters = dto.Parameters
	response.Body.CreatedAt = dto.CreatedAt
	response.Body.UpdatedAt = dto.UpdatedAt
	return response
}

func ScenarioToUpdateResponse(s *scenario.Scenario) *commands.UpdateScenarioResponse {
	dto := buildScenarioDTO(s)
	response := &commands.UpdateScenarioResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Description = dto.Description
	response.Body.Tag = dto.Tag
	response.Body.Icon = dto.Icon
	response.Body.Context = dto.Context
	response.Body.InputData = dto.InputData
	response.Body.Parameters = dto.Parameters
	response.Body.UpdatedAt = dto.UpdatedAt
	return response
}

func ScenarioToDeleteResponse() *commands.DeleteScenarioResponse {
	response := &commands.DeleteScenarioResponse{}
	response.Body.Success = true
	return response
}

func ScenarioToGetResponse(s *scenario.Scenario) *queries.GetScenarioResponse {
	response := &queries.GetScenarioResponse{}
	response.Body = buildScenarioDTO(s)
	return response
}

// ScenarioToListResponse creates a list response with pagination
func ScenarioToListResponse(page, rpp int) func([]*scenario.Scenario) *queries.ListScenariosResponse {
	return func(scenarios []*scenario.Scenario) *queries.ListScenariosResponse {
		response := &queries.ListScenariosResponse{}
		response.Body.Total = len(scenarios)
		response.Body.Page = page
		response.Body.RPP = rpp
		response.Body.Data = MapSlicePtr(scenarios, buildScenarioDTO)
		return response
	}
}

// Mapper instances for handler injection
var (
	ScenarioCreateMapper = CreateMapperFunc[*scenario.Scenario, *commands.CreateScenarioResponse](ScenarioToCreateResponse)
	ScenarioUpdateMapper = UpdateMapperFunc[*scenario.Scenario, *commands.UpdateScenarioResponse](ScenarioToUpdateResponse)
	ScenarioDeleteMapper = DeleteMapperFunc[*commands.DeleteScenarioResponse](ScenarioToDeleteResponse)
	ScenarioGetMapper    = GetMapperFunc[*scenario.Scenario, *queries.GetScenarioResponse](ScenarioToGetResponse)
)

// ScenarioListMapperFactory creates a list mapper with pagination
func ScenarioListMapperFactory(page, rpp int) ListMapperFunc[scenario.Scenario, *queries.ListScenariosResponse] {
	return ListMapperFunc[scenario.Scenario, *queries.ListScenariosResponse](ScenarioToListResponse(page, rpp))
}
