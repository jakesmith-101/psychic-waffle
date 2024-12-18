package main

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"

	"github.com/jakesmith-101/psychic-waffle/api"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/db/mock"
	"github.com/jakesmith-101/psychic-waffle/util"
)

// Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.
type Options struct {
	Debug bool `help:"Enable debug logging" short:"d" default:"true"`
	Port  int  `help:"Port to listen on" short:"p" default:"8080"`
}

func main() {
	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Initilize PgxPool pointer
		db.Open()
		// Ensure SQL tables and basic data exist
		err := mock.MockAll(true)
		if err != nil {
			util.Log(true, "Mock Error: %v", err)
		}

		// Create a new router & API
		router := http.NewServeMux()
		API := humago.New(router, huma.DefaultConfig("Psychic-Waffle API", "1.0.0"))
		// TODO: middleware to add logging of each api call

		// Bind all endpoints to api (login, signup)
		err = api.Endpoints(API)
		if err != nil {
			util.LogError(err)
		}

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
		hooks.OnStop(func() {
			db.PgxPool.Close()
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
