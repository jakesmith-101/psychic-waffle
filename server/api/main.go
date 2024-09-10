package api

import (
	"fmt"
	"path"

	"github.com/danielgtaylor/huma/v2"
)

var apiVer = "v1"
var rootPath = fmt.Sprintf("/api/%s", apiVer)

func BuildPath(route string) string {
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
