package main

import (
	"fmt"
	"log"
	"net/http"

	"parrotflow/internal/container"
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

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Initialize database
		database, err := gorm.Open(sqlite.Open(options.DbPath), &gorm.Config{})
		FailOnError(err, "failed to connect to database")

		// Run migrations
		err = database.AutoMigrate(
			&models.Scenario{},
			&models.ScenarioRun{},
			&models.Tag{},
			&models.Proxy{},
			&models.Agent{},
		)
		FailOnError(err, "failed to migrate database")

		// Initialize application with Wire DI
		app, err := container.InitializeApp(database)
		FailOnError(err, "failed to initialize application")

		// Setup HTTP router and API
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("Parrot Flow API", "1.0.0"))

		// Register all routes
		routes.RegisterAllRoutes(&api, app)

		// Start server
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}
