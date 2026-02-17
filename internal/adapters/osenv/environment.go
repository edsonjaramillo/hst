package osenv

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/edsonjaramillo/hst/internal/config"
)

type Environment struct{}

func (Environment) HistoryFilePath() (string, error) {
	historyFile := strings.TrimSpace(os.Getenv("HISTFILE"))
	if historyFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return config.DefaultHistoryFile(home), nil
	}

	if strings.HasPrefix(historyFile, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		historyFile = filepath.Join(home, strings.TrimPrefix(historyFile, "~"))
	}

	return historyFile, nil
}
