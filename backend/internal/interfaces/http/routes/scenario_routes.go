package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"parrotflow/internal/interfaces/http/handlers"
)

func RegisterScenarioRoutes(api *huma.API, scenarioHandler *handlers.ScenarioHandler) {

	huma.Register(*api, huma.Operation{
		OperationID: "create-scenario",
		Method:      "POST",
		Path:        "/api/scenarios/",
		Summary:     "Create a new scenario",
		Description: "Create a new browser automation scenario",
		Tags:        []string{"scenarios"},
	}, scenarioHandler.CreateScenario)

	huma.Register(*api, huma.Operation{
		OperationID: "get-scenario",
		Method:      "GET",
		Path:        "/api/scenarios/{id}",
		Summary:     "Get a scenario by ID",
		Description: "Retrieve a specific scenario by its ID",
		Tags:        []string{"scenarios"},
	}, scenarioHandler.GetScenario)

	huma.Register(*api, huma.Operation{
		OperationID: "list-scenarios",
		Method:      "GET",
		Path:        "/api/scenarios/",
		Summary:     "List scenarios",
		Description: "Get a list of scenarios with optional filtering",
		Tags:        []string{"scenarios"},
	}, scenarioHandler.ListScenarios)

	huma.Register(*api, huma.Operation{
		OperationID: "update-scenario",
		Method:      "PATCH",
		Path:        "/api/scenarios/{id}",
		Summary:     "Update a scenario",
		Description: "Update an existing scenario",
		Tags:        []string{"scenarios"},
	}, scenarioHandler.UpdateScenario)

	huma.Register(*api, huma.Operation{
		OperationID: "delete-scenario",
		Method:      "DELETE",
		Path:        "/api/scenarios/{id}",
		Summary:     "Delete a scenario",
		Description: "Delete a scenario by its ID",
		Tags:        []string{"scenarios"},
	}, scenarioHandler.DeleteScenario)
}
