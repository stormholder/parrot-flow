package run

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

var (
	apiTag      = "Scenario Run"
	prefix      = "/api/run"
	prefixFmt   = "%s/"
	prefixIdFmt = "%s/{id}"
)

type RunResource struct {
	service *RunService
}

func New(service *RunService) *RunResource {
	return &RunResource{service}
}

func (rs *RunResource) RegisterRoutes(apiBase *huma.API) {
	huma.Register(*apiBase, huma.Operation{
		OperationID: "get-runs",
		Method:      http.MethodGet,
		Path:        fmt.Sprintf(prefixFmt, prefix),
		Summary:     "Get scenario run list",
		Description: "Get a list of scenario runs by criteria",
		Tags:        []string{apiTag},
	}, rs.GetRuns)
	huma.Register(*apiBase, huma.Operation{
		OperationID: "get-run",
		Method:      http.MethodGet,
		Path:        fmt.Sprintf(prefixIdFmt, prefix),
		Summary:     "Get a specific scenario run",
		Description: "Get a specific scenario run by ID",
		Tags:        []string{apiTag},
	}, rs.GetRun)
	huma.Register(*apiBase, huma.Operation{
		OperationID: "handle-run-command",
		Method:      http.MethodPost,
		Path:        fmt.Sprintf(prefixIdFmt, prefix),
		Summary:     "Execute a specific scenario run command",
		Description: "Execute a specific scenario run command by ID",
		Tags:        []string{apiTag},
	}, rs.GetRun)
}

func (rs *RunResource) GetRuns(ctx context.Context, query *RunQuery) (*RunListResponse, error) {
	list, err := rs.service.FindMany(*query)
	if err != nil {
		return nil, err
	}
	resp := &RunListResponse{}
	resp.Body.Pages = list
	return resp, nil
}

func (rs *RunResource) GetRun(ctx context.Context, input *GetRunByIDRequest) (*RunResponse, error) {
	entity, err := rs.service.FindOne(input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("scenario not found")
	}
	resp := &RunResponse{}
	resp.Body.RunEntity = RunEntity{
		RunListItemEntity: RunListItemEntity{
			Model:      entity.Model,
			ScenarioID: entity.ScenarioID,
			Status:     entity.Status,
			StartedAt:  entity.StartedAt,
			FinishedAt: entity.FinishedAt,
			Parameters: entity.Parameters,
		},
	}
	return resp, nil
}

func (rs *RunResource) HandleRunCommand(ctx context.Context, input *RunCommandRequest) (*RunCommandResponse, error) {
	// TODO
	resp := &RunCommandResponse{}
	resp.Body.ID = input.ID
	resp.Body.HandledAt = time.Now()
	resp.Body.Success = true
	return resp, nil
}
