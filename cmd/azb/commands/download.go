package commands

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/azb/pkg/storage"
	"github.com/spf13/cobra"
)

var (
	flagFile string = "file"
)

func buildDownloadCommand() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:     "download",
		Short:   "Download a blob from the upload container",
		Aliases: []string{"down", "get"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fn, err := cmd.Flags().GetString(flagFile)
			if err != nil {
				os.Exit(1)
			}
			err = storage.DownloadFile(fn, args[0])
			if err != nil {
				fmt.Printf("Error while downloading file: %s\r\n", err)
				os.Exit(1)
			}
		},
	}
	downloadCmd.Flags().String(flagFile, "", "Name of the file to download")
	downloadCmd.MarkFlagFilename(flagFile)
	downloadCmd.MarkFlagRequired(flagFile)
	return downloadCmd
}
