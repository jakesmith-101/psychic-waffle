package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"

	"github.com/jakesmith-101/psychic-waffle/db"
)

// Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

var ApiVer = "v1"
var RootPath = fmt.Sprintf("/api/%s", ApiVer)

func main() {
	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Open DB connection
		db.Open()

		// Create a new router & API
		router := http.NewServeMux()
		api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))

		// Register GET /greeting/{name}
		huma.Register(api, huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/greeting/{name}", RootPath),
			Summary:     "Get a greeting",
			Description: "Get a greeting for a person by name.",
			Tags:        []string{"Greetings"},
		}, func(ctx context.Context, input *struct {
			Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
		}) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
