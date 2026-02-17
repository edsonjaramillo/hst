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

type FZFSelector struct{}

func (FZFSelector) SelectMany(ctx context.Context, candidates []string, preselectAll bool) ([]string, error) {
	if len(candidates) == 0 {
		return []string{}, nil
	}

	args := fzfArgs(preselectAll)
	cmd := exec.CommandContext(ctx, "fzf", args...)
	cmd.Stdin = strings.NewReader(strings.Join(candidates, "\n") + "\n")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, &ports.MissingDependencyError{Command: "fzf"}
		}
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			if exitErr.ExitCode() == 1 || exitErr.ExitCode() == 130 {
				return []string{}, nil
			}
		}
		return nil, fmt.Errorf("fzf failed: %w: %s", err, strings.TrimSpace(stderr.String()))
	}

	selected := strings.TrimSpace(stdout.String())
	if selected == "" {
		return []string{}, nil
	}

	return strings.Split(selected, "\n"), nil
}

func fzfArgs(preselectAll bool) []string {
	bindings := []string{"alt-a:select-all", "alt-d:deselect-all"}
	if preselectAll {
		bindings = append([]string{"start:select-all"}, bindings...)
	}

	return []string{"--multi", "--bind", strings.Join(bindings, ",")}
}
