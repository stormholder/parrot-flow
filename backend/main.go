package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"parrotflow/internal/api/scenario"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
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

		scenarioService := scenario.NewScenarioService(scenario.NewStore(database))
		scenarioRoutes := scenario.New(scenarioService)

		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("Parrot Flow API", "1.0.0"))

		huma.Register(api, huma.Operation{
			OperationID: "root",
			Method:      http.MethodGet,
			Path:        "/",
			Summary:     "Root",
			Tags:        []string{"system"},
		}, func(ctx context.Context, i *struct{}) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = "Hello, world!"
			return resp, nil
		})

		scenarioRoutes.RegisterRoutes(&api)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}
