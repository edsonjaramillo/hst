package remove

import (
	"context"

	"edsonjaramillo/hst/internal/common/atuin"
	"edsonjaramillo/hst/internal/common/gum"

	"github.com/urfave/cli/v3"
)

var searchCommand = &cli.Command{
	Name:   "search",
	Usage:  "search for commands to delete",
	Action: searchAction,
}

// searchAction allows the user to search and select commands to delete from atuin history.
func searchAction(_ context.Context, command *cli.Command) error {
	commands := atuin.GetCommandsSortedByFrequency()

	choices := gum.FilterNoLimit(commands)

	atuin.DeleteCommands(choices)

	return nil
}
