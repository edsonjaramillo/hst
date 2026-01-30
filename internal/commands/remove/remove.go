// Package remove provides commands for deleting entries from atuin history.
package remove

import (
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:  "remove",
	Usage: "remove commands from atuin history",
	Commands: []*cli.Command{
		searchCommand,
		errorsCommand,
	},
}
