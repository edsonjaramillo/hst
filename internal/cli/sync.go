package cli

import (
	"github.com/spf13/cobra"

	"github.com/edsonjaramillo/hst/internal/app"
)

func newSyncCmd(uc app.SyncHistoryUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Sync atuin with $HISTFILE",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return uc.Run(cmd.Context())
		},
	}
}
