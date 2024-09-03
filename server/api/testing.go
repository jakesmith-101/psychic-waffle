package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func HealthCheck(api huma.API) {
	// Register GET /greeting/{name}
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        BuildPath("/healthcheck"),
		Summary:     "Get a greeting",
		Description: "Get a greeting.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct{}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = "Hello, world!"
		return resp, nil
	})
}
