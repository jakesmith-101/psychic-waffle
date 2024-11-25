package api

import (
	"context"
	"errors"
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/util"
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

func Endpoints(api huma.API) error {
	err := HealthCheck(api)
	if err != nil {
		return err
	}

	// Auth
	err = AuthEndpoints(api)
	if err != nil {
		return err
	}

	// User
	err = UserEndpoints(api)
	if err != nil {
		return err
	}

	// Role
	err = RoleEndpoints(api)
	if err != nil {
		return err
	}

	// Post
	err = PostEndpoints(api)
	if err != nil {
		return err
	}

	// Comment
	err = CommentEndpoints(api)
	if err != nil {
		return err
	}

	return err
}

type EndpointArgs struct {
	Method  string
	Summary string
	Path    string `required:"true"`
}

var capitals = regexp.MustCompile("[A-Z]")
var prefixReg = regexp.MustCompile(`.*/api\.`) // selects package part of func name (for removal)

func CreateEndpoint[I, O any](api huma.API, op EndpointArgs, handler func(context.Context, *I) (*O, error)) error {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		util.Log(true, "functionName: runtime.Caller: failed")
		return errors.New("functionName: runtime.Caller: failed")
	}

	name := prefixReg.ReplaceAllString(runtime.FuncForPC(counter).Name(), "")
	opID := strings.Trim(capitals.ReplaceAllStringFunc(name, func(m string) string {
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
	util.Log(false, "init: %s", opID)

	return nil
}

// TODO: add authorization header in inputs that require user be logged in: struct {Auth    string `header:"Authorization"`}
