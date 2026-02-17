package ports

import "context"

// HistoryStore wraps Atuin history operations needed by the app layer.
type HistoryStore interface {
	SearchAllWithDuplicates(ctx context.Context) ([]string, error)
	ListHistory(ctx context.Context) ([]string, error)
	DeleteCommands(ctx context.Context, commands []string) error
	DeleteErrorCommands(ctx context.Context) error
}
