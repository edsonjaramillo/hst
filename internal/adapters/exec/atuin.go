package exec

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/edsonjaramillo/hst/internal/ports"
)

type AtuinHistoryStore struct{}

func (AtuinHistoryStore) SearchAllWithDuplicates(ctx context.Context) ([]string, error) {
	return runLines(ctx, "atuin", "search", "--format={command}", "--include-duplicates")
}

func (AtuinHistoryStore) ListHistory(ctx context.Context) ([]string, error) {
	return runLines(ctx, "atuin", "history", "list", "--format", "{command}")
}

func (AtuinHistoryStore) DeleteCommands(ctx context.Context, commands []string) error {
	for _, command := range commands {
		if _, err := runRaw(ctx, "atuin", "search", command, "--delete"); err != nil {
			return err
		}
	}
	return nil
}

func (AtuinHistoryStore) DeleteErrorCommands(ctx context.Context) error {
	_, err := runRaw(ctx, "atuin", "search", "--exclude-exit=0", "--delete-it-all")
	return err
}

func runLines(ctx context.Context, name string, args ...string) ([]string, error) {
	out, err := runRaw(ctx, name, args...)
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimRight(out, "\n")
	if trimmed == "" {
		return []string{}, nil
	}
	return strings.Split(trimmed, "\n"), nil
}

func runRaw(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		return stdout.String(), nil
	}

	if errors.Is(err, exec.ErrNotFound) {
		return "", &ports.MissingDependencyError{Command: name}
	}

	return "", fmt.Errorf("%s %v failed: %w: %s", name, args, err, strings.TrimSpace(stderr.String()))
}
