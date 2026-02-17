package app

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/edsonjaramillo/hst/internal/ports"
)

type fakeHistory struct {
	search []string
	err    error
}

func (f fakeHistory) SearchAllWithDuplicates(context.Context) ([]string, error) {
	return f.search, f.err
}
func (f fakeHistory) ListHistory(context.Context) ([]string, error)  { return nil, nil }
func (f fakeHistory) DeleteCommands(context.Context, []string) error { return nil }
func (f fakeHistory) DeleteErrorCommands(context.Context) error      { return nil }

type fakeSelector struct {
	selected []string
	err      error
}

func (f fakeSelector) SelectMany(context.Context, []string, bool) ([]string, error) {
	return f.selected, f.err
}

func TestRemoveFewerNoCandidates(t *testing.T) {
	uc := RemoveFewerUseCase{
		History:  fakeHistory{search: []string{"a", "a"}},
		Selector: fakeSelector{selected: []string{"a"}},
	}
	err := uc.Run(context.Background(), 0)
	if !IsCode(err, CodeNoCandidates) {
		t.Fatalf("expected CodeNoCandidates, got %v", err)
	}
}

func TestRemoveFewerMissingDependency(t *testing.T) {
	uc := RemoveFewerUseCase{
		History:  fakeHistory{search: []string{"a"}},
		Selector: fakeSelector{err: &ports.MissingDependencyError{Command: "fzf"}},
	}
	err := uc.Run(context.Background(), 1)
	if !IsCode(err, CodeDependencyMissing) {
		t.Fatalf("expected CodeDependencyMissing, got %v", err)
	}
}

func TestRemoveFewerInvalidArg(t *testing.T) {
	uc := RemoveFewerUseCase{}
	err := uc.Run(context.Background(), -1)
	if !IsCode(err, CodeInvalidArgument) {
		t.Fatalf("expected CodeInvalidArgument, got %v", err)
	}
}

func TestRemoveFewerNoSelection(t *testing.T) {
	uc := RemoveFewerUseCase{
		History:  fakeHistory{search: []string{"a"}},
		Selector: fakeSelector{selected: nil},
	}
	err := uc.Run(context.Background(), 1)
	if !IsCode(err, CodeNoSelection) {
		t.Fatalf("expected CodeNoSelection, got %v", err)
	}
}

func TestRemoveFewerAlphabeticalCandidates(t *testing.T) {
	selector := &selectorSpy{selected: []string{"a"}}
	uc := RemoveFewerUseCase{
		History:  fakeHistory{search: []string{"b", "a", "a", "c", "c", "c"}},
		Selector: selector,
	}

	err := uc.Run(context.Background(), 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantCandidates := []string{"a", "b"}
	if !reflect.DeepEqual(selector.candidates, wantCandidates) {
		t.Fatalf("got candidates %v want %v", selector.candidates, wantCandidates)
	}
	if !selector.preselectAll {
		t.Fatalf("expected preselectAll to be true")
	}
}

func TestWrapDependencyPassesOperational(t *testing.T) {
	err := wrapDependency("oops", errors.New("x"))
	if !IsCode(err, CodeOperational) {
		t.Fatalf("expected CodeOperational, got %v", err)
	}
}
