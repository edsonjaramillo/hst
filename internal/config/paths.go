package config

import "path/filepath"

// DefaultHistoryFile returns the conventional shell history file path.
func DefaultHistoryFile(home string) string {
	return filepath.Join(home, ".zsh_history")
}
