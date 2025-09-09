package scenario

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

var (
	apiTag = "Scenario"
)

type ScenarioResource struct {
	service *ScenarioService
}

func New(service *ScenarioService) *ScenarioResource {
	return &ScenarioResource{service}
}

func (rs *ScenarioResource) RegisterRoutes(apiBase *huma.API) {
	prefix := "/api/scenarios"
	huma.Register(*apiBase, huma.Operation{
		OperationID: "get-scenarios",
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/", prefix),
		Summary:     "Get scenarios",
		Description: "Get a list of scenarios by criteria",
		Tags:        []string{apiTag},
	}, rs.GetScenarios)
	huma.Register(*apiBase, huma.Operation{
		OperationID:   "create-scenario",
		Method:        http.MethodPost,
		Path:          fmt.Sprintf("%s/", prefix),
		Summary:       "Create new scenario",
		Description:   "Create an empty new scenario",
		Tags:          []string{apiTag},
		DefaultStatus: http.StatusCreated,
	}, rs.CreateScenario)
	huma.Register(*apiBase, huma.Operation{
		OperationID: "get-scenario",
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/{id}", prefix),
		Summary:     "Get a specific scenario",
		Description: "Get a specific scenario by ID",
		Tags:        []string{apiTag},
	}, rs.GetScenario)
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
	var payload ScenarioPayload
	if err := json.Unmarshal([]byte(scenario.Payload), &payload); err != nil {
		return nil, err
	}
	resp := &ScenarioResponse{}
	resp.Body.ScenarioBase = scenario.ScenarioBase
	resp.Body.Payload = payload
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
