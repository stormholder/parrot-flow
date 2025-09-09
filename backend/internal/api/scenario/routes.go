package scenario

import (
	"context"
	"fmt"
	"net/http"

	"parrotflow/internal/api"

	"github.com/danielgtaylor/huma/v2"
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
		Tags:        []string{"scenario"},
	}, rs.GetScenarios)
	huma.Register(*apiBase, huma.Operation{
		OperationID:   "create-scenario",
		Method:        http.MethodPost,
		Path:          fmt.Sprintf("%s/", prefix),
		Summary:       "Create new scenario",
		Description:   "Create an empty new scenario",
		Tags:          []string{"scenario"},
		DefaultStatus: http.StatusCreated,
	}, rs.CreateScenario)
}

func (rs *ScenarioResource) GetScenarios(ctx context.Context, query *ScenarioQuery) (*api.Pages, error) {
	resp, err := rs.service.FindMany(*query)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
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
