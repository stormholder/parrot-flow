package main

import (
	"fmt"
	"log"
	"net/http"
	"parrotflow/internal/infrastructure/events"
	"parrotflow/internal/infrastructure/persistence"
	"parrotflow/internal/interfaces/http/routes"
	"parrotflow/internal/models"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Options struct {
	Port   int    `help:"Port to listen on" short:"p" default:"8888"`
	DbPath string `help:"Database file path" short:"d" default:"store.db"`
}

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		database, err := gorm.Open(sqlite.Open(options.DbPath), &gorm.Config{})
		FailOnError(err, "failed to connect to database")

		err = database.AutoMigrate(
			&models.Scenario{},
			&models.ScenarioRun{},
			&models.Tag{},
			&models.Proxy{},
			&models.Agent{},
		)
		FailOnError(err, "failed to migrate database")

		scenarioRepository := persistence.NewScenarioRepository(database)
		runRepository := persistence.NewRunRepository(database)
		tagRepository := persistence.NewTagRepository(database)
		proxyRepository := persistence.NewProxyRepository(database)
		agentRepository := persistence.NewAgentRepository(database)

		eventBus := events.NewAsyncEventBus()
		eventBus.Subscribe(events.NewScenarioCreatedHandler())
		eventBus.Subscribe(events.NewScenarioUpdatedHandler())
		eventBus.Subscribe(events.NewScenarioDeletedHandler())
		eventBus.Subscribe(events.NewRunCreatedHandler())
		eventBus.Subscribe(events.NewRunStartedHandler())
		eventBus.Subscribe(events.NewRunCompletedHandler())
		eventBus.Subscribe(events.NewRunFailedHandler())

		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("Parrot Flow API", "1.0.0"))

		routes.RegisterSystemRoutes(&api)
		routes.RegisterScenarioRoutes(&api, scenarioRepository, eventBus)
		routes.RegisterRunRoutes(&api, runRepository, scenarioRepository, eventBus)
		routes.RegisterTagRoutes(&api, tagRepository, eventBus)
		routes.RegisterProxyRoutes(&api, proxyRepository, eventBus)
		routes.RegisterAgentRoutes(&api, agentRepository, eventBus)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}
