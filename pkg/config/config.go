package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	UseCliAuth         bool
	StorageAccountName string
	StorageAccountKey  string
}

func Load() *Config {
	return &Config{
		StorageAccountName: viper.GetString("storage_account_name"),
		UseCliAuth:         viper.GetBool("auth.cli"),
		StorageAccountKey:  viper.GetString("auth.key"),
	}
}

func (c *Config) Save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	f := path.Join(home, ".azb.yaml")
	viper.Set("storage_account_name", c.StorageAccountName)
	viper.Set("auth.cli", c.UseCliAuth)
	viper.Set("auth.key", c.StorageAccountKey)
	viper.SetConfigType("yaml")
	return viper.WriteConfigAs(f)
}

func (c *Config) IsConfigured() bool {
	return len(c.StorageAccountName) > 0
}
