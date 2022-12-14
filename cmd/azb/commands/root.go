package commands

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BuildRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "azb",
		Short: "Dead simple personal backup CLI leveraging Azure storage account",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(buildInitCommand())
	rootCmd.AddCommand(buildInfoCommand())
	rootCmd.AddCommand(buildUploadCommand())
	rootCmd.AddCommand(buildListCommand())
	rootCmd.AddCommand(buildDownloadCommand())
	rootCmd.AddCommand(buildDeleteCommand())

	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfiguration)
}

func initConfiguration() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	defaultCfg := path.Join(home, ".azb.yaml")

	viper.SetConfigFile(defaultCfg)
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("AZB")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

}
