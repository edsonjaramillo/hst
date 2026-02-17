package cli

import (
	"github.com/spf13/cobra"

	"github.com/edsonjaramillo/hst/internal/app"
)

func newRemoveErrorsCmd(uc app.RemoveErrorsUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "errors",
		Short: "deletes commands that resulted in errors",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return uc.Run(cmd.Context())
		},
	}
}
