package routes

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func RegisterSystemRoutes(api *huma.API) {
	huma.Register(*api, huma.Operation{
		OperationID: "root",
		Method:      "GET",
		Path:        "/",
		Summary:     "Root",
		Description: "Root endpoint for the Parrot Flow API",
		Tags:        []string{"system"},
	}, func(ctx context.Context, i *struct{}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = "Hello, world!"
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "health",
		Method:      "GET",
		Path:        "/health",
		Summary:     "Health Check",
		Description: "Check the health status of the API",
		Tags:        []string{"system"},
	}, func(ctx context.Context, i *struct{}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = "OK"
		return resp, nil
	})
}
