// Package gum provides wrappers around the gum interactive command-line tools.
package gum

import (
	"edsonjaramillo/hst/internal/common/shell"
)

// Choose presents a list of choices to the user and returns the selected ones.
func Choose(choices []string) []string {
	parts := append([]string{"choose"}, choices...)
	return shell.CmdInteractive("gum", parts...)
}

// Filter filters the given choices using gum's interactive filter.
func Filter(choices []string) []string {
	parts := append([]string{"filter"}, choices...)
	return shell.CmdInteractive("gum", parts...)
}

// FilterNoLimit filters choices without limiting the number of results displayed.
func FilterNoLimit(choices []string) []string {
	parts := append([]string{"filter", "--no-limit"}, choices...)
	return shell.CmdInteractive("gum", parts...)
}
