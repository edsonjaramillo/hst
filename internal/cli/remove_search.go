package cli

import (
	"github.com/spf13/cobra"

	"github.com/edsonjaramillo/hst/internal/app"
)

func newRemoveSearchCmd(uc app.RemoveSearchUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "search for commands to delete",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return uc.Run(cmd.Context())
		},
	}
}
