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

func HealthCheck(api huma.API) error {
	// Register GET /healthcheck/{name}
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/healthcheck/{name}",
		Summary: "Get a greeting.",
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"John" doc:"Any name, defaults to 'world'"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		name := input.Name
		if name == "" {
			name = "world"
		}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", name)
		fmt.Fprintf(os.Stderr, "Healthy: %s\n", name)
		return resp, nil
	})
}
