package ports

import "fmt"

// MissingDependencyError reports a missing external command dependency.
type MissingDependencyError struct {
	Command string
}

func (e *MissingDependencyError) Error() string {
	return fmt.Sprintf("missing dependency: %s", e.Command)
}
