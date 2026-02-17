package app

import (
	"context"

	"github.com/edsonjaramillo/hst/internal/domain"
	"github.com/edsonjaramillo/hst/internal/ports"
)

type RemoveFewerUseCase struct {
	History  ports.HistoryStore
	Selector ports.Selector
}

func (uc RemoveFewerUseCase) Run(ctx context.Context, fewer int) error {
	if fewer < 0 {
		return NewError(CodeInvalidArgument, "fewer must be greater than or equal to 0", nil)
	}

	commands, err := uc.History.SearchAllWithDuplicates(ctx)
	if err != nil {
		return wrapDependency("failed to read commands from atuin", err)
	}

	candidates := domain.CommandsWithMaxFrequency(commands, fewer)
	if len(candidates) == 0 {
		return NewError(CodeNoCandidates, "No commands found with fewer than the specified occurences.", nil)
	}

	selected, err := uc.Selector.SelectMany(ctx, candidates, true)
	if err != nil {
		return wrapDependency("failed to run interactive selector", err)
	}
	if len(selected) == 0 {
		return NewError(CodeNoSelection, "No commands selected for deletion.", nil)
	}

	if err := uc.History.DeleteCommands(ctx, selected); err != nil {
		return wrapDependency("failed to delete selected commands", err)
	}

	return nil
}
