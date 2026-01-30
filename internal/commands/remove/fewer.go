package remove

import (
	"context"

	"edsonjaramillo/hst/internal/common/atuin"
	"edsonjaramillo/hst/internal/common/gum"
	"edsonjaramillo/hst/internal/common/shell"

	"github.com/urfave/cli/v3"
)

var fewerCommand = &cli.Command{
	Name:   "fewer",
	Usage:  "delete commands that occur fewer than certain occurences",
	Action: fewerAction,
	Arguments: []cli.Argument{
		fewerArg,
	},
}

var fewerArg = &cli.IntArg{
	Name:      "fewer",
	Value:     1,
	UsageText: "number of occurences to delete",
}

// fewerAction allows the user to search and select commands to delete from atuin history.
func fewerAction(_ context.Context, command *cli.Command) error {
	commands := atuin.GetAllCommandsUsed()
	fewer := command.IntArg("fewer")

	frequencyCommands := atuin.GetCommandWithMaxFrequency(commands, fewer)
	if len(frequencyCommands) == 0 {
		shell.Exit("No commands found with fewer than the specified occurences.")
	}

	choices := gum.FilterNoLimit(frequencyCommands, true)
	if len(choices) == 0 {
		shell.Exit("No commands selected for deletion.")
	}

	atuin.DeleteCommands(choices)

	return nil
}
