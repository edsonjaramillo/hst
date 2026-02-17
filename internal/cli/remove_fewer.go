package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/edsonjaramillo/hst/internal/app"
)

func newRemoveFewerCmd(uc app.RemoveFewerUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "fewer [fewer]",
		Short: "delete commands that occur fewer than certain occurences",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return cobra.MaximumNArgs(1)(cmd, args)
			}
			if len(args) == 1 {
				_, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fewer := 1
			if len(args) == 1 {
				parsed, err := strconv.Atoi(args[0])
				if err != nil {
					return app.NewError(app.CodeInvalidArgument, "fewer must be an integer", err)
				}
				fewer = parsed
			}
			return uc.Run(cmd.Context(), fewer)
		},
	}
}
