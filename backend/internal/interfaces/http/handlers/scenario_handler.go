package handlers

import (
	"context"

	command "parrotflow/internal/application/command/scenario"
	query "parrotflow/internal/application/query/scenario"
	"parrotflow/internal/domain/scenario"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

type ScenarioHandler struct {
	createCommandHandler *command.CreateScenarioCommandHandler
	updateCommandHandler *command.UpdateScenarioCommandHandler
	deleteCommandHandler *command.DeleteScenarioCommandHandler
	getQueryHandler      *query.GetScenarioQueryHandler
	listQueryHandler     *query.ListScenariosQueryHandler

	// Mappers - using functional types
	createMapper mappers.CreateMapperFunc[*scenario.Scenario, *commands.CreateScenarioResponse]
	updateMapper mappers.UpdateMapperFunc[*scenario.Scenario, *commands.UpdateScenarioResponse]
	deleteMapper mappers.DeleteMapperFunc[*commands.DeleteScenarioResponse]
	getMapper    mappers.GetMapperFunc[*scenario.Scenario, *queries.GetScenarioResponse]
}

func NewScenarioHandler(
	createCommandHandler *command.CreateScenarioCommandHandler,
	updateCommandHandler *command.UpdateScenarioCommandHandler,
	deleteCommandHandler *command.DeleteScenarioCommandHandler,
	getQueryHandler *query.GetScenarioQueryHandler,
	listQueryHandler *query.ListScenariosQueryHandler,
) *ScenarioHandler {
	return &ScenarioHandler{
		createCommandHandler: createCommandHandler,
		updateCommandHandler: updateCommandHandler,
		deleteCommandHandler: deleteCommandHandler,
		getQueryHandler:      getQueryHandler,
		listQueryHandler:     listQueryHandler,
		createMapper:         mappers.ScenarioCreateMapper,
		updateMapper:         mappers.ScenarioUpdateMapper,
		deleteMapper:         mappers.ScenarioDeleteMapper,
		getMapper:            mappers.ScenarioGetMapper,
	}
}

func (h *ScenarioHandler) CreateScenario(ctx context.Context, req *commands.CreateScenarioRequest) (*commands.CreateScenarioResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.CreateScenarioRequest) (command.CreateScenarioCommand, error) {
			return command.CreateScenarioCommand{
				Name:        r.Body.Name,
				Description: r.Body.Description,
				Tag:         r.Body.Tag,
				Icon:        r.Body.Icon,
			}, nil
		},
		CommandHandlerFunc[command.CreateScenarioCommand, *scenario.Scenario](h.createCommandHandler.Handle),
		h.createMapper,
	)
}

func (h *ScenarioHandler) GetScenario(ctx context.Context, req *queries.GetScenarioRequest) (*queries.GetScenarioResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetScenarioRequest) (query.GetScenarioQuery, error) {
			scenarioID, err := scenario.NewScenarioID(r.ID)
			if err != nil {
				return query.GetScenarioQuery{}, err
			}
			return query.GetScenarioQuery{ID: scenarioID}, nil
		},
		QueryHandlerFunc[query.GetScenarioQuery, *scenario.Scenario](h.getQueryHandler.Handle),
		h.getMapper,
	)
}

func (h *ScenarioHandler) ListScenarios(ctx context.Context, req *queries.ListScenariosRequest) (*queries.ListScenariosResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.ListScenariosRequest) (query.ListScenariosQuery, error) {
			limit := r.RPP
			offset := (r.Page - 1) * r.RPP
			return query.ListScenariosQuery{Criteria: scenario.SearchCriteria{
				Limit:  limit,
				Offset: offset,
			}}, nil
		},
		QueryHandlerFunc[query.ListScenariosQuery, []*scenario.Scenario](h.listQueryHandler.Handle),
		mappers.ScenarioListMapperFactory(req.Page, req.RPP),
	)
}

func (h *ScenarioHandler) UpdateScenario(ctx context.Context, req *commands.UpdateScenarioRequest) (*commands.UpdateScenarioResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.UpdateScenarioRequest) (command.UpdateScenarioCommand, error) {
			scenarioID, err := scenario.NewScenarioID(r.ID)
			if err != nil {
				return command.UpdateScenarioCommand{}, err
			}

			cmd := command.UpdateScenarioCommand{
				ID:          scenarioID,
				Name:        r.Body.Name,
				Description: r.Body.Description,
				Tag:         r.Body.Tag,
				Icon:        r.Body.Icon,
			}

			// Map value objects if provided
			if r.Body.Context != nil {
				ctx := mappers.MapContextFromDTO(*r.Body.Context)
				cmd.Context = &ctx
			}
			if r.Body.InputData != nil {
				inputData := mappers.MapInputDataFromDTO(*r.Body.InputData)
				cmd.InputData = &inputData
			}
			if r.Body.Parameters != nil {
				params := mappers.MapParametersFromDTO(*r.Body.Parameters)
				cmd.Parameters = &params
			}

			return cmd, nil
		},
		CommandHandlerFunc[command.UpdateScenarioCommand, *scenario.Scenario](h.updateCommandHandler.Handle),
		h.updateMapper,
	)
}

func (h *ScenarioHandler) DeleteScenario(ctx context.Context, req *commands.DeleteScenarioRequest) (*commands.DeleteScenarioResponse, error) {
	return HandleSimpleCommand(
		ctx,
		req,
		func(r *commands.DeleteScenarioRequest) (command.DeleteScenarioCommand, error) {
			scenarioID, err := scenario.NewScenarioID(r.ID)
			if err != nil {
				return command.DeleteScenarioCommand{}, err
			}
			return command.DeleteScenarioCommand{ID: scenarioID}, nil
		},
		SimpleCommandHandlerFunc[command.DeleteScenarioCommand](h.deleteCommandHandler.Handle),
		h.deleteMapper.Map,
	)
}
