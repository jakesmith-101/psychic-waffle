package api

import (
	"fmt"
	"path"
)

var apiVer = "v1"
var rootPath = fmt.Sprintf("/api/%s", apiVer)

func BuildPath(route string) string {
	return path.Join(rootPath, route)
}
