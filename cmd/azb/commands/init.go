package commands

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/azb/pkg/config"
	"github.com/spf13/cobra"
)

func buildInitCommand() *cobra.Command {
	flagCliAuth := "use-cli-auth"
	flagAccountName := "storage-account-name"
	flagAccountKey := "storage-account-key"

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize your instance of azb",
		Run: func(cmd *cobra.Command, args []string) {
			sa, err := cmd.Flags().GetString(flagAccountName)
			cobra.CheckErr(err)
			cli, err := cmd.Flags().GetBool(flagCliAuth)
			cobra.CheckErr(err)
			key, err := cmd.Flags().GetString(flagAccountKey)
			cobra.CheckErr(err)

			cfg := config.Load()
			cfg.StorageAccountName = sa
			cfg.UseCliAuth = cli
			if cfg.UseCliAuth {
				cfg.StorageAccountKey = ""
			} else if len(key) > 0 {
				cfg.StorageAccountKey = key
				cfg.UseCliAuth = false
			}
			if err := cfg.Save(); err != nil {
				fmt.Printf("Error while updating config: %v", err)
				os.Exit(1)
			}

		},
	}

	initCmd.Flags().String(flagAccountName, "", "Storage Account Name")
	initCmd.MarkFlagRequired(flagAccountName)

	initCmd.Flags().String(flagAccountKey, "", "Client ID used for authentication")
	initCmd.Flags().Bool(flagCliAuth, false, "Re-Use Azure CLI authentication")
	return initCmd
}
