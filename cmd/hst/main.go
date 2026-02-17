package main

import (
	"os"

	"github.com/edsonjaramillo/hst/internal/cli"
	"github.com/edsonjaramillo/hst/internal/cli/output"
)

func main() {
	err := cli.Execute()
	if err == nil {
		return
	}

	if !output.IsAlreadyReportedFailure(err) {
		_ = output.PrintRootError(os.Stderr, err)
	}
	os.Exit(output.ExitCode(err))
}
