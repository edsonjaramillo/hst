package app

import (
	"errors"
	"fmt"
)

type ErrorCode string

const (
	CodeInvalidArgument     ErrorCode = "invalid_argument"
	CodeDependencyMissing   ErrorCode = "dependency_missing"
	CodeNoCandidates        ErrorCode = "no_candidates"
	CodeNoSelection         ErrorCode = "no_selection"
	CodeHistoryFileNotFound ErrorCode = "history_file_not_found"
	CodeOperational         ErrorCode = "operational"
)

// Error represents a typed application error for consistent exit-code mapping.
type Error struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func NewError(code ErrorCode, message string, err error) error {
	return &Error{Code: code, Message: message, Err: err}
}

func IsCode(err error, code ErrorCode) bool {
	var appErr *Error
	if !errors.As(err, &appErr) {
		return false
	}
	return appErr.Code == code
}
