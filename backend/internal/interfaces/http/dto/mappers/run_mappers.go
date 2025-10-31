package mappers

import (
	"parrotflow/internal/domain/run"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

func buildRunDTO(r *run.Run) queries.RunListItem {
	startedAt := ""
	if r.StartedAt != nil {
		startedAt = FormatTimestamp(r.StartedAt.Time())
	}
	finishedAt := ""
	if r.FinishedAt != nil {
		finishedAt = FormatTimestamp(r.FinishedAt.Time())
	}

	return queries.RunListItem{
		ID:         r.Id.String(),
		ScenarioID: r.ScenarioID.String(),
		Status:     r.Status.String(),
		Parameters: r.Parameters,
		StartedAt:  &startedAt,
		FinishedAt: &finishedAt,
		CreatedAt:  FormatTimestamp(r.CreatedAt.Time()),
	}
}

// Mapper functions using functional approach

func RunToCreateResponse(r *run.Run) *commands.CreateRunResponse {
	dto := buildRunDTO(r)
	response := &commands.CreateRunResponse{}
	response.Body.ID = dto.ID
	response.Body.ScenarioID = dto.ScenarioID
	response.Body.Status = dto.Status
	response.Body.CreatedAt = dto.CreatedAt
	return response
}

func RunToStartResponse(r *run.Run) *commands.StartRunResponse {
	dto := buildRunDTO(r)
	response := &commands.StartRunResponse{}
	response.Body.ID = dto.ID
	response.Body.Status = dto.Status
	if dto.StartedAt != nil {
		response.Body.StartedAt = *dto.StartedAt
	}
	return response
}

func RunToGetResponse(r *run.Run) *queries.GetRunResponse {
	dto := buildRunDTO(r)
	response := &queries.GetRunResponse{}
	response.Body.ID = dto.ID
	response.Body.ScenarioID = dto.ScenarioID
	response.Body.Status = dto.Status
	response.Body.Parameters = dto.Parameters
	response.Body.StartedAt = dto.StartedAt
	response.Body.FinishedAt = dto.FinishedAt
	response.Body.CreatedAt = dto.CreatedAt
	return response
}

func RunToListResponse(runs []*run.Run) *queries.ListRunsResponse {
	response := &queries.ListRunsResponse{}
	response.Body.Data = MapSlicePtr(runs, buildRunDTO)
	response.Body.Total = len(runs)
	response.Body.Page = 1
	response.Body.RPP = len(runs)
	return response
}

// Mapper instances for handler injection
var (
	RunCreateMapper = CreateMapperFunc[*run.Run, *commands.CreateRunResponse](RunToCreateResponse)
	RunStartMapper  = CreateMapperFunc[*run.Run, *commands.StartRunResponse](RunToStartResponse)
	RunGetMapper    = GetMapperFunc[*run.Run, *queries.GetRunResponse](RunToGetResponse)
	RunListMapper   = ListMapperFunc[run.Run, *queries.ListRunsResponse](RunToListResponse)
)
