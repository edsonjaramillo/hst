package domain

import (
	"reflect"
	"testing"
)

func TestUniqueSortedAlphabetical(t *testing.T) {
	input := []string{"z", "a", "m", "a", "z", "b"}
	got := UniqueSortedAlphabetical(input)
	want := []string{"a", "b", "m", "z"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestCommandsWithMaxFrequency(t *testing.T) {
	input := []string{"a", "a", "b", "c", "c", "c"}
	got := CommandsWithMaxFrequency(input, 2)
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestUniqueNonEmptyCommands(t *testing.T) {
	input := []string{"  ls  ", "", "ls", "git\nstatus", "git status"}
	got := UniqueNonEmptyCommands(input)
	want := []string{"ls", "git status"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
