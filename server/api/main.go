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
var rootPathLength = len(rootPath)

func BuildPath(route string) string {
	if len(route) >= rootPathLength && route[:rootPathLength] == rootPath {
		return route
	}
	return path.Join(rootPath, route)
}

func AllEndpoints(api huma.API) {
	HealthCheck(api)
	AuthEndpoints(api)
	UserEndpoints(api)
	PostEndpoints(api)
	CommentEndpoints(api)
}

func AuthEndpoints(api huma.API) {
	Signup(api)
	Login(api)
}

func UserEndpoints(api huma.API) {
	GetUser(api)
	UpdateUser(api)
}

func PostEndpoints(api huma.API) {
	GetPosts(api)
	GetPopularPosts(api)
}

func CommentEndpoints(api huma.API) {

}

type EndpointArgs struct {
	Name    string
	Method  string
	Summary string
	Path    string
}

var reg = regexp.MustCompile("[A-Z]")

func CreateEndpoint[I, O any](api huma.API, op EndpointArgs, handler func(context.Context, *I) (*O, error)) {
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
