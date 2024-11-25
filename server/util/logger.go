package util

import (
	"fmt"
	"os"
)

func Log(isError bool, format string, a ...any) {
	output := fmt.Sprintf(format, a...)
	if isError {
		fmt.Fprintf(os.Stderr, "%s\n", output)
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", output)
	}
}

func LogError(err error) {
	Log(true, "%v", err)
}
