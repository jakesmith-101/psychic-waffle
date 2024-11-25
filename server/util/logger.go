package util

import (
	"fmt"
	"os"
)

func Log(t string, format string, a ...any) {
	output := fmt.Sprintf(format, a)
	if t == "error" {
		fmt.Fprintf(os.Stderr, "%s\n", output)
	} else if t == "output" {
		fmt.Fprintf(os.Stdout, "%s\n", output)
	}
}

func LogError(err error) {
	Log("error", fmt.Sprintf("%v", err))
}
