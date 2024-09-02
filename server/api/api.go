package api

import (
	"fmt"
)

var ApiVer = "v1"
var RootPath = fmt.Sprintf("/api/%s", ApiVer)

func BuildPath(path string) string {
	return fmt.Sprintf("%s%s", RootPath, path)
}
