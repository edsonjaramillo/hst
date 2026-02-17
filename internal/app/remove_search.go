package app

import (
	"context"
	"errors"

	"github.com/edsonjaramillo/hst/internal/domain"
	"github.com/edsonjaramillo/hst/internal/ports"
)

type RemoveSearchUseCase struct {
	History  ports.HistoryStore
	Selector ports.Selector
}

func (uc RemoveSearchUseCase) Run(ctx context.Context) error {
	commands, err := uc.History.SearchAllWithDuplicates(ctx)
	if err != nil {
		return wrapDependency("failed to read commands from atuin", err)
	}

	candidates := domain.UniqueSortedAlphabetical(commands)
	if len(candidates) == 0 {
		return NewError(CodeNoCandidates, "No commands found to delete.", nil)
	}

	selected, err := uc.Selector.SelectMany(ctx, candidates, false)
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

func wrapDependency(message string, err error) error {
	var missing *ports.MissingDependencyError
	if errors.As(err, &missing) {
		return NewError(CodeDependencyMissing, missing.Error(), err)
	}
	return NewError(CodeOperational, message, err)
}
