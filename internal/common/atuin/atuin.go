// Package atuin provides functions for interacting with the atuin shell history manager.
package atuin

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"edsonjaramillo/hst/internal/common/shell"
)

// GetAllCommandsUsed returns all commands ever used from atuin history.
func GetAllCommandsUsed() []string {
	allCommands := shell.CmdInteractive("atuin", "search", "--format={command}", "--include-duplicates")
	return allCommands
}

// GetCommandsSortedByFrequency returns unique commands sorted by how frequently they were used (least to most).
func GetCommandsSortedByFrequency() []string {
	allCommands := GetAllCommandsUsed()

	// Count how many times each command appears
	frequency := make(map[string]int)
	for _, cmd := range allCommands {
		frequency[cmd]++
	}

	// Extract unique commands into a slice
	uniqueCommands := make([]string, 0, len(frequency))
	for cmd := range frequency {
		uniqueCommands = append(uniqueCommands, cmd)
	}

	// Sort commands by frequency (least common first, most common last)
	sort.Slice(uniqueCommands, func(i, j int) bool {
		return frequency[uniqueCommands[i]] < frequency[uniqueCommands[j]]
	})

	return uniqueCommands
}

func GetCommandWithMaxFrequency(commands []string, fewerThan int) []string {
	frequency := make(map[string]int)
	for _, cmd := range commands {
		frequency[cmd]++
	}

	commandsToDelete := []string{}
	for cmd, freq := range frequency {
		if freq <= fewerThan {
			commandsToDelete = append(commandsToDelete, cmd)
		}
	}

	return commandsToDelete
}

// DeleteCommands deletes the given commands from atuin history.
func DeleteCommands(commands []string) {
	for _, command := range commands {
		err := shell.Cmd("atuin", "search", command, "--delete").Run()
		if err != nil {
			shell.Exit("atuin: could not delete " + command + ": ")
		}
	}
}

// DeleteErrorCommands deletes all commands from atuin history that resulted in errors (non-zero exit codes).
func DeleteErrorCommands() {
	err := shell.Cmd("atuin", "search", "--exclude-exit=0", "--delete-it-all").Run()
	if err != nil {
		shell.Exit("atuin: could not delete error commands: ")
	}
}

// FileExists checks if a file exists at the given filepath.
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

// SyncHistory syncs atuin history to the file specified by $HISTFILE.
func SyncHistory() {
	zshHistoryFile, err := getHistoryFilePath()
	if err != nil {
		shell.Exit("atuin: could not determine history file path: " + err.Error())
	}

	if !FileExists(zshHistoryFile) {
		shell.Exit("atuin: zsh history file does not exist: " + zshHistoryFile)
	}

	output := shell.CmdInteractive("atuin", "history", "list", "--format", "{command}")
	uniqueCommands := uniqueNonEmptyCommands(output)

	file, err := os.Create(zshHistoryFile)
	if err != nil {
		shell.Exit("atuin: could not create history file: " + err.Error())
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	writer := bufio.NewWriter(file)

	for _, line := range uniqueCommands {
		_, writeErr := writer.WriteString(line + "\n")
		if writeErr != nil {
			shell.Exit("atuin: could not write to history file.")
		}
	}

	if flushErr := writer.Flush(); flushErr != nil {
		shell.Exit("atuin: could not flush history file.")
	}
}

func getHistoryFilePath() (string, error) {
	historyFile := strings.TrimSpace(os.Getenv("HISTFILE"))
	if historyFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".zsh_history"), nil
	}

	if strings.HasPrefix(historyFile, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		historyFile = filepath.Join(home, strings.TrimPrefix(historyFile, "~"))
	}

	return historyFile, nil
}

func uniqueNonEmptyCommands(commands []string) []string {
	seen := make(map[string]struct{}, len(commands))
	unique := make([]string, 0, len(commands))

	for _, cmd := range commands {
		cleaned := strings.TrimSpace(strings.ReplaceAll(cmd, "\n", " "))
		if cleaned == "" {
			continue
		}
		if _, exists := seen[cleaned]; exists {
			continue
		}
		seen[cleaned] = struct{}{}
		unique = append(unique, cleaned)
	}

	return unique
}
