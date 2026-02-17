package cli

import (
	"github.com/spf13/cobra"

	"github.com/edsonjaramillo/hst/internal/adapters/exec"
	"github.com/edsonjaramillo/hst/internal/adapters/fs"
	"github.com/edsonjaramillo/hst/internal/adapters/osenv"
	"github.com/edsonjaramillo/hst/internal/app"
)

// Version is injected at build time.
var Version = "dev"

func Execute() error {
	return newRootCmd().Execute()
}

func newRootCmd() *cobra.Command {
	history := exec.AtuinHistoryStore{}
	selector := exec.FZFSelector{}
	historyFile := fs.HistoryFile{}
	environment := osenv.Environment{}

	removeSearchUC := app.RemoveSearchUseCase{History: history, Selector: selector}
	removeErrorsUC := app.RemoveErrorsUseCase{History: history}
	removeFewerUC := app.RemoveFewerUseCase{History: history, Selector: selector}
	syncUC := app.SyncHistoryUseCase{History: history, HistoryFile: historyFile, Env: environment}

	cmd := &cobra.Command{
		Use:           "hst",
		Short:         "hst is a cli helper for atuin",
		Long:          "hst manages shell history maintenance workflows backed by atuin.",
		Version:       Version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(newRemoveCmd(removeSearchUC, removeErrorsUC, removeFewerUC))
	cmd.AddCommand(newSyncCmd(syncUC))
	cmd.AddCommand(newCompletionCmd(cmd))

	return cmd
}
