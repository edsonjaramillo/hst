package app

import (
	"context"
	"reflect"
	"testing"

	"github.com/edsonjaramillo/hst/internal/ports"
)

type searchHistoryStub struct {
	search  []string
	err     error
	deleted []string
}

func (s *searchHistoryStub) SearchAllWithDuplicates(context.Context) ([]string, error) {
	return s.search, s.err
}

func (s *searchHistoryStub) ListHistory(context.Context) ([]string, error) { return nil, nil }
func (s *searchHistoryStub) DeleteCommands(_ context.Context, commands []string) error {
	s.deleted = commands
	return nil
}
func (s *searchHistoryStub) DeleteErrorCommands(context.Context) error { return nil }

type selectorSpy struct {
	selected     []string
	err          error
	candidates   []string
	preselectAll bool
}

func (s *selectorSpy) SelectMany(_ context.Context, candidates []string, preselectAll bool) ([]string, error) {
	s.candidates = candidates
	s.preselectAll = preselectAll
	return s.selected, s.err
}

func TestRemoveSearchAlphabeticalCandidates(t *testing.T) {
	history := &searchHistoryStub{search: []string{"z", "a", "a", "m", "z", "b"}}
	selector := &selectorSpy{selected: []string{"a"}}
	uc := RemoveSearchUseCase{
		History:  history,
		Selector: selector,
	}

	err := uc.Run(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantCandidates := []string{"a", "b", "m", "z"}
	if !reflect.DeepEqual(selector.candidates, wantCandidates) {
		t.Fatalf("got candidates %v want %v", selector.candidates, wantCandidates)
	}
	if selector.preselectAll {
		t.Fatalf("expected preselectAll to be false")
	}
	if !reflect.DeepEqual(history.deleted, []string{"a"}) {
		t.Fatalf("got deleted %v want %v", history.deleted, []string{"a"})
	}
}

func TestRemoveSearchNoCandidates(t *testing.T) {
	uc := RemoveSearchUseCase{
		History:  &searchHistoryStub{search: nil},
		Selector: &selectorSpy{selected: []string{"a"}},
	}

	err := uc.Run(context.Background())
	if !IsCode(err, CodeNoCandidates) {
		t.Fatalf("expected CodeNoCandidates, got %v", err)
	}
}

func TestRemoveSearchMissingDependency(t *testing.T) {
	uc := RemoveSearchUseCase{
		History:  &searchHistoryStub{search: []string{"a"}},
		Selector: &selectorSpy{err: &ports.MissingDependencyError{Command: "fzf"}},
	}

	err := uc.Run(context.Background())
	if !IsCode(err, CodeDependencyMissing) {
		t.Fatalf("expected CodeDependencyMissing, got %v", err)
	}
}

func TestRemoveSearchNoSelection(t *testing.T) {
	uc := RemoveSearchUseCase{
		History:  &searchHistoryStub{search: []string{"a"}},
		Selector: &selectorSpy{selected: nil},
	}

	err := uc.Run(context.Background())
	if !IsCode(err, CodeNoSelection) {
		t.Fatalf("expected CodeNoSelection, got %v", err)
	}
}
