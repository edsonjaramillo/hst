package ports

import "context"

// HistoryFile handles history file existence checks and atomic writes.
type HistoryFile interface {
	PathExists(ctx context.Context, path string) (bool, error)
	WriteLinesAtomic(ctx context.Context, path string, lines []string) error
}
