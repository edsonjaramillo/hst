package cli

import (
	"github.com/spf13/cobra"

	"github.com/edsonjaramillo/hst/internal/app"
)

func newRemoveCmd(searchUC app.RemoveSearchUseCase, errorsUC app.RemoveErrorsUseCase, fewerUC app.RemoveFewerUseCase) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove commands from atuin history",
	}

	cmd.AddCommand(newRemoveSearchCmd(searchUC))
	cmd.AddCommand(newRemoveErrorsCmd(errorsUC))
	cmd.AddCommand(newRemoveFewerCmd(fewerUC))

	return cmd
}
