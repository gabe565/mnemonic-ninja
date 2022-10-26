package config

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

func Parse() (err error) {
	flag.Parse()
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		if err := viper.BindPFlag(strings.ReplaceAll(f.Name, "-", "."), f); err != nil {
			return
		}
	})
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/mnemonic-ninja/")
	viper.AddConfigPath("$HOME/.mnemonic-ninja")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MNEMONIC_NINJA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("Fatal error reading config file: %w \n", err)
		}
	}
	return nil
}
