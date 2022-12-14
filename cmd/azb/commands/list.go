package commands

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/azb/pkg/storage"
	"github.com/spf13/cobra"
)

func buildListCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Short:   "List all blobs stored in your backup container",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			err := storage.ListFiles()
			if err != nil {
				fmt.Printf("Error while listing files: %s\r\n", err)
				os.Exit(1)
			}
		},
	}

	return listCmd
}
