package commands

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/azb/pkg/storage"
	"github.com/spf13/cobra"
)

func buildDeleteCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete a blob from the upload container",
		Aliases: []string{"del", "rm"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args[0]) == 0 {
				os.Exit(1)
			}
			err := storage.DeleteBlob(args[0])
			if err != nil {
				fmt.Printf("Error while deleting blob: %s\r\n", err)
				os.Exit(1)
			}
		},
	}

	return listCmd
}
