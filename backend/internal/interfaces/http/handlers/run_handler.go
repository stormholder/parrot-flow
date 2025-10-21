package handlers

import (
	"context"

	command "parrotflow/internal/application/command/run"
	query "parrotflow/internal/application/query/run"
	"parrotflow/internal/domain/run"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

type RunHandler struct {
	createCommandHandler *command.CreateRunCommandHandler
	startCommandHandler  *command.StartRunCommandHandler
	getQueryHandler      *query.GetRunQueryHandler
	listQueryHandler     *query.ListRunsQueryHandler

	createMapper mappers.RunCreateMapper
	startMapper  mappers.RunStartMapper
	getMapper    mappers.RunGetMapper
	listMapper   mappers.RunListMapper
}

func NewRunHandler(
	createCommandHandler *command.CreateRunCommandHandler,
	startCommandHandler *command.StartRunCommandHandler,
	getQueryHandler *query.GetRunQueryHandler,
	listQueryHandler *query.ListRunsQueryHandler,
) *RunHandler {
	return &RunHandler{
		createCommandHandler: createCommandHandler,
		startCommandHandler:  startCommandHandler,
		getQueryHandler:      getQueryHandler,
		listQueryHandler:     listQueryHandler,
		createMapper:         mappers.RunCreateMapper{},
		startMapper:          mappers.RunStartMapper{},
		getMapper:            mappers.RunGetMapper{},
		listMapper:           mappers.RunListMapper{},
	}
}

func (h *RunHandler) CreateRun(ctx context.Context, req *commands.CreateRunRequest) (*commands.CreateRunResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.CreateRunRequest) (command.CreateRunCommand, error) {
			scenarioID, err := scenario.NewScenarioID(r.Body.ScenarioID)
			if err != nil {
				return command.CreateRunCommand{}, err
			}
			return command.CreateRunCommand{ScenarioID: scenarioID}, nil
		},
		CommandHandlerFunc[command.CreateRunCommand, *run.Run](h.createCommandHandler.Handle),
		h.createMapper,
	)
}

func (h *RunHandler) StartRun(ctx context.Context, req *commands.StartRunRequest) (*commands.StartRunResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.StartRunRequest) (command.StartRunCommand, error) {
			runID, err := run.NewRunID(r.ID)
			if err != nil {
				return command.StartRunCommand{}, err
			}
			return command.StartRunCommand{RunID: runID}, nil
		},
		CommandHandlerFunc[command.StartRunCommand, *run.Run](h.startCommandHandler.Handle),
		h.startMapper,
	)
}

func (h *RunHandler) GetRun(ctx context.Context, req *queries.GetRunRequest) (*queries.GetRunResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetRunRequest) (query.GetRunQuery, error) {
			runID, err := run.NewRunID(r.ID)
			if err != nil {
				return query.GetRunQuery{}, err
			}
			return query.GetRunQuery{ID: runID}, nil
		},
		QueryHandlerFunc[query.GetRunQuery, *run.Run](h.getQueryHandler.Handle),
		h.getMapper,
	)
}

func (h *RunHandler) ListRuns(ctx context.Context, req *queries.ListRunsRequest) (*queries.ListRunsResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.ListRunsRequest) (query.ListRunsQuery, error) {
			criteria := run.NewSearchCriteria()
			if r.ScenarioID != "" {
				scenarioID, err := scenario.NewScenarioID(r.ScenarioID)
				if err != nil {
					return query.ListRunsQuery{}, err
				}
				criteria.ScenarioID = scenarioID
			}
			return query.ListRunsQuery{Criteria: criteria}, nil
		},
		QueryHandlerFunc[query.ListRunsQuery, []*run.Run](h.listQueryHandler.Handle),
		h.listMapper,
	)
}
