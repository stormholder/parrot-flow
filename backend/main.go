package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"

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
	// port := "8888"
	// dbFile := "store.db"

	// if fromEnv := os.Getenv("PORT"); fromEnv != "" {
	// 	port = fromEnv
	// }

	// if dbPath := os.Getenv("DATABASE"); dbPath != "" {
	// 	dbFile = dbPath
	// }

	// database, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	// FailOnError(err, "failed to connect to database")

	// log.Printf("Starting up on http://localhost:%s", port)

	// scenarioService := scenario.NewScenarioService(scenario.NewStore(database))

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("Parrot Flow API", "1.0.0"))

		huma.Register(api, huma.Operation{
			OperationID: "root",
			Method:      http.MethodGet,
			Path:        "/",
			Summary:     "Root",
			Tags:        []string{"system"},
		}, func(ctx context.Context, i *interface{}) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello!")
			return resp, nil
		})
	})

	cli.Run()
}
