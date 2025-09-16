package scenario

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"parrotflow/internal/models"

	"github.com/danielgtaylor/huma/v2"
)

var (
	apiTag      = "Scenario"
	prefix      = "/api/scenarios"
	prefixFmt   = "%s/"
	prefixIdFmt = "%s/{id}"
)

type ScenarioResource struct {
	service *ScenarioService
}

func New(service *ScenarioService) *ScenarioResource {
	return &ScenarioResource{service}
}

func (rs *ScenarioResource) RegisterRoutes(apiBase *huma.API) {
	huma.Register(*apiBase, huma.Operation{
		OperationID: "get-scenarios",
		Method:      http.MethodGet,
		Path:        fmt.Sprintf(prefixFmt, prefix),
		Summary:     "Get scenarios",
		Description: "Get a list of scenarios by criteria",
		Tags:        []string{apiTag},
	}, rs.GetScenarios)
	huma.Register(*apiBase, huma.Operation{
		OperationID:   "create-scenario",
		Method:        http.MethodPost,
		Path:          fmt.Sprintf(prefixFmt, prefix),
		Summary:       "Create new scenario",
		Description:   "Create an empty new scenario",
		Tags:          []string{apiTag},
		DefaultStatus: http.StatusCreated,
	}, rs.CreateScenario)
	huma.Register(*apiBase, huma.Operation{
		OperationID: "get-scenario",
		Method:      http.MethodGet,
		Path:        fmt.Sprintf(prefixIdFmt, prefix),
		Summary:     "Get a specific scenario",
		Description: "Get a specific scenario by ID",
		Tags:        []string{apiTag},
	}, rs.GetScenario)
	huma.Register(*apiBase, huma.Operation{
		OperationID: "update-scenario",
		Method:      http.MethodPatch,
		Path:        fmt.Sprintf(prefixIdFmt, prefix),
		Summary:     "Update a specific scenario",
		Description: "Update a specific scenario by ID",
		Tags:        []string{apiTag},
	}, rs.UpdateScenario)
	huma.Register(*apiBase, huma.Operation{
		OperationID: "delete-scenario",
		Method:      http.MethodDelete,
		Path:        fmt.Sprintf(prefixIdFmt, prefix),
		Summary:     "Delete a specific scenario",
		Description: "Delete a specific scenario by ID",
		Tags:        []string{apiTag},
	}, rs.DeleteScenario)
}

func (rs *ScenarioResource) GetScenarios(ctx context.Context, query *ScenarioQuery) (*ScenarioListResponse, error) {
	list, err := rs.service.FindMany(*query)
	if err != nil {
		return nil, err
	}
	resp := &ScenarioListResponse{}
	resp.Body.Pages = list
	return resp, nil
}

func (rs *ScenarioResource) GetScenario(ctx context.Context, input *GetScenarioByIDRequest) (*ScenarioResponse, error) {
	scenario, err := rs.service.FindOne(input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("scenario not found")
	}
	var context ScenarioContext
	if err := json.Unmarshal([]byte(scenario.Context), &context); err != nil {
		return nil, err
	}
	resp := &ScenarioResponse{}
	resp.Body.ScenarioBase = scenario.ScenarioBase
	resp.Body.Context = context
	return resp, nil
}

func (rs *ScenarioResource) CreateScenario(ctx context.Context, i *struct{}) (*ScenarioCreateResponse, error) {
	newScenario, err := rs.service.Create()
	if err != nil {
		return nil, err
	}
	resp := &ScenarioCreateResponse{}
	resp.Body.Scenario = newScenario
	return resp, nil
}

func (rs *ScenarioResource) UpdateScenario(ctx context.Context, input *ScenarioPatchRequest) (*models.Scenario, error) {
	_, err := rs.service.FindOne(input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("scenario not found")
	}
	updated, err := rs.service.Update(input.ID, input.Body.ScenarioPatch)
	if err != nil {
		return nil, huma.Error500InternalServerError("")
	}
	return updated, nil
}

func (rs *ScenarioResource) DeleteScenario(ctx context.Context, input *GetScenarioByIDRequest) (*struct{}, error) {
	return rs.service.Delete(uint(input.ID))
}
