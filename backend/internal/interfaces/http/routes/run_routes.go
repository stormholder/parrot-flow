package routes

import (
	"net/http"
	commandRun "parrotflow/internal/application/command/run"
	queryRun "parrotflow/internal/application/query/run"
	"parrotflow/internal/infrastructure/events"
	"parrotflow/internal/infrastructure/persistence"
	"parrotflow/internal/interfaces/http/handlers"

	"github.com/danielgtaylor/huma/v2"
)

var (
	apiTag  = []string{"ScenarioRun"}
	apiPath = "/api/runs"
)

func RegisterRunRoutes(
	api *huma.API,
	runRepository *persistence.RunRepository,
	scenarioRepository *persistence.ScenarioRepository,
	eventBus *events.AsyncEventBus,
) {
	runHandler := handlers.NewRunHandler(
		commandRun.NewCreateRunCommandHandler(runRepository, scenarioRepository, eventBus),
		commandRun.NewStartRunCommandHandler(runRepository, eventBus),
		queryRun.NewGetRunQueryHandler(runRepository),
		queryRun.NewListRunsQueryHandler(runRepository),
	)

	huma.Register(*api, huma.Operation{
		OperationID: "create-run",
		Method:      http.MethodPost,
		Path:        apiPath,
		Summary:     "Create a new run",
		Description: "Create a new scenario run",
		Tags:        apiTag,
	}, runHandler.CreateRun)

	huma.Register(*api, huma.Operation{
		OperationID: "get-run",
		Method:      http.MethodGet,
		Path:        "/api/runs/{id}",
		Summary:     "Get a run by ID",
		Description: "Retrieve a specific run by its ID",
		Tags:        apiTag,
	}, runHandler.GetRun)

	huma.Register(*api, huma.Operation{
		OperationID: "list-runs",
		Method:      http.MethodGet,
		Path:        apiPath,
		Summary:     "List runs",
		Description: "Get a list of runs with optional filtering",
		Tags:        apiTag,
	}, runHandler.ListRuns)

	huma.Register(*api, huma.Operation{
		OperationID: "start-run",
		Method:      http.MethodPost,
		Path:        "/api/runs/{id}/start",
		Summary:     "Start a run",
		Description: "Start execution of a run",
		Tags:        apiTag,
	}, runHandler.StartRun)
}
