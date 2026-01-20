// Package shell provides utilities for executing shell commands.
package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Cmd creates a new command without interactive I/O.
func Cmd(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	return cmd
}

// CmdInteractive creates a new command with stdin/stdout connected to the terminal and returns the output.
func CmdInteractive(name string, args ...string) []string {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		Exit("shell: could not execute interactive command: " + err.Error())
	}

	return outputCleaner(output, "\n")
}

// Exit prints a red error message and exits with status 1
func Exit(message string) {
	fmt.Println("\033[31m" + message + "\033[0m")
	os.Exit(1)
}

// outputCleaner splits command output by delimiter and removes empty strings
func outputCleaner(output []byte, delimiter string) []string {
	splits := strings.Split(string(output), delimiter)
	lines := []string{}
	for _, line := range splits {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
