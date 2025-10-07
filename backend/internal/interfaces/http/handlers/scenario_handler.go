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
	}
}

func (h *ScenarioHandler) CreateScenario(ctx context.Context, req *commands.CreateScenarioRequest) (*commands.CreateScenarioResponse, error) {
	cmd := command.CreateScenarioCommand{
		Name:        req.Body.Name,
		Description: req.Body.Description,
		Tag:         req.Body.Tag,
		Icon:        req.Body.Icon,
	}

	scenario, err := h.createCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	response := mappers.ToCreateScenarioResponse(scenario)

	return response, nil
}

func (h *ScenarioHandler) GetScenario(ctx context.Context, req *queries.GetScenarioRequest) (*queries.GetScenarioResponse, error) {
	scenarioID, err := scenario.NewScenarioID(req.ID)
	if err != nil {
		return nil, err
	}

	query := query.GetScenarioQuery{
		ID: scenarioID,
	}

	scenario, err := h.getQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	response := mappers.ToGetScenarioResponse(scenario)

	return response, nil
}

func (h *ScenarioHandler) ListScenarios(ctx context.Context, req *queries.ListScenariosRequest) (*queries.ListScenariosResponse, error) {
	criteria := scenario.NewSearchCriteria()
	if req.Name != "" {
		criteria = criteria.WithName(req.Name)
	}
	if req.Tag != "" {
		criteria = criteria.WithTag(req.Tag)
	}
	if req.Page > 0 {
		criteria = criteria.WithPagination(req.RPP, (req.Page-1)*req.RPP)
	}

	query := query.ListScenariosQuery{
		Criteria: criteria,
	}

	scenarios, err := h.listQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	response := mappers.ToListScenarioResponse(scenarios, req.Page, req.RPP)

	return response, nil
}

func (h *ScenarioHandler) UpdateScenario(ctx context.Context, req *commands.UpdateScenarioRequest) (*commands.UpdateScenarioResponse, error) {
	scenarioID, err := scenario.NewScenarioID(req.ID)
	if err != nil {
		return nil, err
	}

	cmd := command.UpdateScenarioCommand{
		ID:          scenarioID,
		Name:        req.Body.Name,
		Description: req.Body.Description,
		Tag:         req.Body.Tag,
		Icon:        req.Body.Icon,
	}

	scenario, err := h.updateCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	response := mappers.ToScenarioUpdateResponse(scenario)

	return response, nil
}

func (h *ScenarioHandler) DeleteScenario(ctx context.Context, req *commands.DeleteScenarioRequest) (*commands.DeleteScenarioResponse, error) {
	scenarioID, err := scenario.NewScenarioID(req.ID)
	if err != nil {
		return nil, err
	}

	cmd := command.DeleteScenarioCommand{
		ID: scenarioID,
	}

	err = h.deleteCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	response := &commands.DeleteScenarioResponse{}
	response.Body.Success = true

	return response, nil
}
