package commands

import (
	"fmt"

	"github.com/ThorstenHans/azb/pkg/config"
	"github.com/spf13/cobra"
)

func buildInfoCommand() *cobra.Command {
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Print contextual information about azb",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load()
			fmt.Printf("azb is using Azure Storage Account '%s'\r\n", cfg.StorageAccountName)
			if cfg.UseCliAuth {
				fmt.Printf("azb is relying on Azure CLI authentication\r\n")
			}
		},
	}

	return infoCmd
}
