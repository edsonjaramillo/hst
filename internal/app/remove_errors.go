package app

import (
	"context"

	"github.com/edsonjaramillo/hst/internal/ports"
)

type RemoveErrorsUseCase struct {
	History ports.HistoryStore
}

func (uc RemoveErrorsUseCase) Run(ctx context.Context) error {
	if err := uc.History.DeleteErrorCommands(ctx); err != nil {
		return wrapDependency("failed to delete error commands", err)
	}
	return nil
}
