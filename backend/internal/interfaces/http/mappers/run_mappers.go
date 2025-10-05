package mappers

import (
	"parrotflow/internal/domain/run"
	"parrotflow/internal/interfaces/http/commands"
	"parrotflow/internal/interfaces/http/queries"
)

func ToGetRunResponse(run *run.Run) *queries.GetRunResponse {
	response := &queries.GetRunResponse{}
	response.Body.ID = run.Id.String()
	response.Body.ScenarioID = run.ScenarioID.String()
	response.Body.Status = run.Status.String()
	response.Body.Parameters = run.Parameters
	response.Body.CreatedAt = run.CreatedAt.Time().Format("2006-01-02T15:04:05Z")
	response.Body.UpdatedAt = run.UpdatedAt.Time().Format("2006-01-02T15:04:05Z")

	if run.StartedAt != nil {
		startedAt := run.StartedAt.Time().Format("2006-01-02T15:04:05Z")
		response.Body.StartedAt = &startedAt
	}
	if run.FinishedAt != nil {
		finishedAt := run.FinishedAt.Time().Format("2006-01-02T15:04:05Z")
		response.Body.FinishedAt = &finishedAt
	}

	return response
}

func ToCreateRunResponse(run *run.Run) *commands.CreateRunResponse {
	response := &commands.CreateRunResponse{}
	response.Body.ID = run.Id.String()
	response.Body.ScenarioID = run.ScenarioID.String()
	response.Body.Status = run.Status.String()
	response.Body.Parameters = run.Parameters
	response.Body.CreatedAt = run.CreatedAt.Time().Format("2006-01-02T15:04:05Z")

	return response
}

func toRunListItem(run *run.Run) *queries.RunListItem {
	item := &queries.RunListItem{}

	item.ID = run.Id.String()
	item.ScenarioID = run.ScenarioID.String()
	item.Status = run.Status.String()
	item.Parameters = run.Parameters
	item.CreatedAt = run.CreatedAt.Time().Format("2006-01-02T15:04:05Z")

	if run.StartedAt != nil {
		startedAt := run.StartedAt.Time().Format("2006-01-02T15:04:05Z")
		item.StartedAt = &startedAt
	}
	if run.FinishedAt != nil {
		finishedAt := run.FinishedAt.Time().Format("2006-01-02T15:04:05Z")
		item.FinishedAt = &finishedAt
	}

	return item
}

func ToListRunResponse(runs []*run.Run, page int, rpp int) *queries.ListRunsResponse {
	response := &queries.ListRunsResponse{}
	response.Body.Total = len(runs)
	response.Body.Page = page
	response.Body.RPP = rpp

	response.Body.Data = make([]queries.RunListItem, len(runs))

	for i, run := range runs {
		response.Body.Data[i] = *toRunListItem(run)
	}

	return response
}
