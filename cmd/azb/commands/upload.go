package commands

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/azb/pkg/storage"
	"github.com/spf13/cobra"
)

func buildUploadCommand() *cobra.Command {
	uploadCmd := &cobra.Command{
		Use:     "upload",
		Short:   "Upload a file to your upload container in Azure Storage",
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{"up"},
		Run: func(cmd *cobra.Command, args []string) {
			err := storage.UploadFile(args[0])
			if err != nil {
				fmt.Printf("Error while uploading file: %s\r\n", err)
				os.Exit(1)
			}
		},
	}
	return uploadCmd
}
