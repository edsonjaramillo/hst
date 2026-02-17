package app

import (
	"context"
	"fmt"

	"github.com/edsonjaramillo/hst/internal/domain"
	"github.com/edsonjaramillo/hst/internal/ports"
)

type SyncHistoryUseCase struct {
	History     ports.HistoryStore
	HistoryFile ports.HistoryFile
	Env         ports.Environment
}

func (uc SyncHistoryUseCase) Run(ctx context.Context) error {
	path, err := uc.Env.HistoryFilePath()
	if err != nil {
		return NewError(CodeOperational, "could not resolve history file path", err)
	}

	exists, err := uc.HistoryFile.PathExists(ctx, path)
	if err != nil {
		return NewError(CodeOperational, "could not check history file path", err)
	}
	if !exists {
		return NewError(CodeHistoryFileNotFound, fmt.Sprintf("history file does not exist: %s", path), nil)
	}

	entries, err := uc.History.ListHistory(ctx)
	if err != nil {
		return wrapDependency("failed to list atuin history", err)
	}

	clean := domain.UniqueNonEmptyCommands(entries)
	if err := uc.HistoryFile.WriteLinesAtomic(ctx, path, clean); err != nil {
		return NewError(CodeOperational, "could not write history file", err)
	}

	return nil
}
