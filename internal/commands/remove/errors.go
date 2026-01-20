package remove

import (
	"context"

	"edsonjaramillo/hst/internal/common/atuin"

	"github.com/urfave/cli/v3"
)

var errorsCommand = &cli.Command{
	Name:   "errors",
	Usage:  "deletes commands that resulted in errors",
	Action: errorsAction,
}

// errorsAction deletes all commands that resulted in errors (non-zero exit codes).
func errorsAction(_ context.Context, command *cli.Command) error {
	atuin.DeleteErrorCommands()
	return nil
}
