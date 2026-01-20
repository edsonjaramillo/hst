// Package sync provides commands for syncing history with external sources.
package sync

import (
	"context"

	"edsonjaramillo/hst/internal/common/atuin"

	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:   "sync",
	Usage:  "Sync atuin with $HISTFILE",
	Action: action,
}

// action syncs atuin history to the $HISTFILE and removes error commands.
func action(_ context.Context, command *cli.Command) error {
	atuin.DeleteErrorCommands()
	return nil
}
