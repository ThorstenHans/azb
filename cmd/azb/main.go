package main

import (
	"os"

	"github.com/ThorstenHans/azb/cmd/azb/commands"
)

func main() {
	rootCmd := commands.BuildRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
