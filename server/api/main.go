package api

import (
	"context"
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

var apiVer = "v1"
var rootPath = fmt.Sprintf("/api/%s", apiVer)

func BuildPath(route string) string {
	// TODO: test route does not already contain rootPath
	return path.Join(rootPath, route)
}

func AllEndpoints(api huma.API) {
	HealthCheck(api)
	AuthEndpoints(api)
	UserEndpoints(api)
}

func AuthEndpoints(api huma.API) {
	Signup(api)
	Login(api)
}

func UserEndpoints(api huma.API) {
	GetUser(api)
	UpdateUser(api)
}

type EndpointArgs struct {
	Name    string
	Method  string
	Summary string
	Path    string
}

func CreateEndpoint[I, O any](api huma.API, op EndpointArgs, handler func(context.Context, *I) (*O, error)) {
	reg := regexp.MustCompile("[A-Z]")
	words := reg.Split(op.Name, -1)
	opID := strings.ToLower(strings.Join(words, "-"))

	huma.Register(api, huma.Operation{
		OperationID: opID,
		Method:      op.Method,
		Path:        BuildPath(op.Path),
		Summary:     op.Summary,
		Description: op.Summary,
		Tags:        []string{op.Name},
	}, handler)
}
