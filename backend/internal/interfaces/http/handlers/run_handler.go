package handlers

import (
	"context"
	commandRun "parrotflow/internal/application/command/run"
	queryRun "parrotflow/internal/application/query/run"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

type RunHandler struct {
	createCommandHandler *commandRun.CreateRunCommandHandler
	startCommandHandler  *commandRun.StartRunCommandHandler
	getQueryHandler      *queryRun.GetRunQueryHandler
	listQueryHandler     *queryRun.ListRunsQueryHandler
}

func NewRunHandler(
	createCommandHandler *commandRun.CreateRunCommandHandler,
	startCommandHandler *commandRun.StartRunCommandHandler,
	getQueryHandler *queryRun.GetRunQueryHandler,
	listQueryHandler *queryRun.ListRunsQueryHandler,
) *RunHandler {
	return &RunHandler{
		createCommandHandler: createCommandHandler,
		startCommandHandler:  startCommandHandler,
		getQueryHandler:      getQueryHandler,
		listQueryHandler:     listQueryHandler,
	}
}

func (h *RunHandler) CreateRun(ctx context.Context, req *commands.CreateRunRequest) (*commands.CreateRunResponse, error) {
	scenarioID, err := scenario.NewScenarioID(req.Body.ScenarioID)
	if err != nil {
		return nil, err
	}

	cmd := commandRun.CreateRunCommand{
		ScenarioID: scenarioID,
		Parameters: req.Body.Parameters,
	}

	run, err := h.createCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	response := mappers.ToCreateRunResponse(run)

	return response, nil
}

func (h *RunHandler) GetRun(ctx context.Context, req *queries.GetRunRequest) (*queries.GetRunResponse, error) {
	runID, err := run.NewRunID(req.ID)
	if err != nil {
		return nil, err
	}

	query := queryRun.GetRunQuery{
		ID: runID,
	}

	run, err := h.getQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	response := mappers.ToGetRunResponse(run)

	return response, nil
}

func (h *RunHandler) ListRuns(ctx context.Context, req *queries.ListRunsRequest) (*queries.ListRunsResponse, error) {
	criteria := run.NewSearchCriteria()
	if req.ScenarioID != "" {
		scenarioID, err := scenario.NewScenarioID(req.ScenarioID)
		if err != nil {
			return nil, err
		}
		criteria = criteria.WithScenarioID(scenarioID)
	}
	if req.Status != "" {
		criteria = criteria.WithStatus(req.Status)
	}
	if req.Page > 0 {
		criteria = criteria.WithPagination(req.RPP, (req.Page-1)*req.RPP)
	}

	query := queryRun.ListRunsQuery{
		Criteria: criteria,
	}

	runs, err := h.listQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	response := mappers.ToListRunResponse(runs, req.Page, req.RPP)

	return response, nil
}

func (h *RunHandler) StartRun(ctx context.Context, req *commands.StartRunRequest) (*commands.StartRunResponse, error) {
	runID, err := run.NewRunID(req.ID)
	if err != nil {
		return nil, err
	}

	cmd := commandRun.StartRunCommand{
		RunID: runID,
	}

	run, err := h.startCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	response := &commands.StartRunResponse{}
	response.Body.ID = run.Id.String()
	response.Body.Status = run.Status.String()
	if run.StartedAt != nil {
		response.Body.StartedAt = run.StartedAt.Time().Format("2006-01-02T15:04:05Z")
	}

	return response, nil
}
