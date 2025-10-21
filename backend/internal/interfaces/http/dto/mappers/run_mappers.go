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

type RunCreateMapper struct{}

func (m RunCreateMapper) Map(r *run.Run) *commands.CreateRunResponse {
	dto := buildRunDTO(r)
	response := &commands.CreateRunResponse{}
	response.Body.ID = dto.ID
	response.Body.ScenarioID = dto.ScenarioID
	response.Body.Status = dto.Status
	response.Body.CreatedAt = dto.CreatedAt
	return response
}

type RunStartMapper struct{}

func (m RunStartMapper) Map(r *run.Run) *commands.StartRunResponse {
	dto := buildRunDTO(r)
	response := &commands.StartRunResponse{}
	response.Body.ID = dto.ID
	response.Body.Status = dto.Status
	if dto.StartedAt != nil {
		response.Body.StartedAt = *dto.StartedAt
	}
	return response
}

type RunGetMapper struct{}

func (m RunGetMapper) Map(r *run.Run) *queries.GetRunResponse {
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

type RunListMapper struct{}

func (m RunListMapper) Map(runs []*run.Run) *queries.ListRunsResponse {
	response := &queries.ListRunsResponse{}
	response.Body.Data = MapSlicePtr(runs, buildRunDTO)
	response.Body.Total = len(runs)
	response.Body.Page = 1
	response.Body.RPP = len(runs)
	return response
}

func ToCreateRunResponse(r *run.Run) *commands.CreateRunResponse {
	return RunCreateMapper{}.Map(r)
}

func ToStartRunResponse(r *run.Run) *commands.StartRunResponse {
	return RunStartMapper{}.Map(r)
}

func ToGetRunResponse(r *run.Run) *queries.GetRunResponse {
	return RunGetMapper{}.Map(r)
}

func ToListRunsResponse(runs []*run.Run) *queries.ListRunsResponse {
	return RunListMapper{}.Map(runs)
}
