package api

import (
	"fmt"
)

var apiVer = "v1"
var rootPath = fmt.Sprintf("/api/%s", apiVer)

func BuildPath(path string) string {
	return fmt.Sprintf("%s%s", rootPath, path)
}
