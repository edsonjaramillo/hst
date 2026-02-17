package output

import (
	"errors"
	"fmt"
	"io"

	"github.com/edsonjaramillo/hst/internal/app"
)

var errAlreadyReportedFailure = errors.New("command failed with details already reported")

func AlreadyReportedFailure() error {
	return errAlreadyReportedFailure
}

func IsAlreadyReportedFailure(err error) bool {
	return errors.Is(err, errAlreadyReportedFailure)
}

func PrintRootError(w io.Writer, err error) error {
	_, writeErr := fmt.Fprintf(w, "Error: %v\n", err)
	return writeErr
}

func ExitCode(err error) int {
	if err == nil {
		return 0
	}
	if IsAlreadyReportedFailure(err) {
		return 1
	}

	switch {
	case app.IsCode(err, app.CodeInvalidArgument):
		return 2
	case app.IsCode(err, app.CodeDependencyMissing):
		return 2
	case app.IsCode(err, app.CodeNoCandidates):
		return 2
	case app.IsCode(err, app.CodeNoSelection):
		return 2
	case app.IsCode(err, app.CodeHistoryFileNotFound):
		return 2
	default:
		return 1
	}
}
