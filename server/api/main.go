package api

import (
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
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
	GetComments(api)
	GetPopularComments(api)
}

type EndpointArgs struct {
	Method  string
	Summary string
	Path    string `required:"true"`
}

var reg = regexp.MustCompile("[A-Z]")

func CreateEndpoint[I, O any](api huma.API, op EndpointArgs, handler func(context.Context, *I) (*O, error)) {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		fmt.Fprintf(os.Stderr, "functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	name := runtime.FuncForPC(counter).Name()
	opID := strings.Trim(reg.ReplaceAllStringFunc(name, func(m string) string {
		return fmt.Sprint("-", strings.ToLower(m))
	}), "-")

	huma.Register(api, huma.Operation{
		OperationID: opID,
		Method:      op.Method,
		Path:        BuildPath(op.Path),
		Summary:     op.Summary,
		Description: op.Summary,
		Tags:        []string{name},
	}, handler)
	fmt.Fprintf(os.Stderr, "init: %s", opID)
}
