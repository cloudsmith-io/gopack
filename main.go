package main

import (
	"os"

	"github.com/cloudsmith-io/gopack/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
