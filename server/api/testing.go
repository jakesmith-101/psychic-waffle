package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
		Path:        BuildPath("/healthcheck/{name}"),
		Summary:     "Get a greeting",
		Description: "Get a greeting.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"John" doc:"Any name, defaults to 'world'"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		name := input.Name
		if name == "" {
			name = "world"
		}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", name)
		fmt.Fprintf(os.Stderr, "Healthy: %s", name)
		return resp, nil
	})
}
