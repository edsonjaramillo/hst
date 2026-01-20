// Package main is the entry point for the CLI application
package main

import (
	"context"
	"log"
	"os"

	"edsonjaramillo/hst/internal/commands/remove"
	"edsonjaramillo/hst/internal/commands/sync"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "hst",
		Version: "0.1.0",
		Usage:   `hst is a cli helper for atuin`,
		Commands: []*cli.Command{
			remove.Command,
			sync.Command,
		},
		HideHelpCommand: true,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
