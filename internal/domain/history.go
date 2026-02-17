package domain

import (
	"sort"
	"strings"
)

// UniqueSortedAlphabetical returns unique commands sorted alphabetically.
func UniqueSortedAlphabetical(commands []string) []string {
	freq := Frequency(commands)
	unique := make([]string, 0, len(freq))
	for command := range freq {
		unique = append(unique, command)
	}
	sort.Strings(unique)
	return unique
}

// CommandsWithMaxFrequency returns unique commands with frequency <= max.
func CommandsWithMaxFrequency(commands []string, max int) []string {
	freq := Frequency(commands)
	filtered := make([]string, 0, len(freq))
	for command, count := range freq {
		if count <= max {
			filtered = append(filtered, command)
		}
	}
	sort.Strings(filtered)
	return filtered
}

// Frequency counts command occurrences.
func Frequency(commands []string) map[string]int {
	freq := make(map[string]int, len(commands))
	for _, command := range commands {
		if command == "" {
			continue
		}
		freq[command]++
	}
	return freq
}

// UniqueNonEmptyCommands trims entries, removes blank lines, and deduplicates while keeping order.
func UniqueNonEmptyCommands(commands []string) []string {
	seen := make(map[string]struct{}, len(commands))
	unique := make([]string, 0, len(commands))

	for _, command := range commands {
		clean := strings.TrimSpace(strings.ReplaceAll(command, "\n", " "))
		if clean == "" {
			continue
		}
		if _, exists := seen[clean]; exists {
			continue
		}
		seen[clean] = struct{}{}
		unique = append(unique, clean)
	}

	return unique
}
