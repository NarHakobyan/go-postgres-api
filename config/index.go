package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	configFile, err := filepath.Abs("./config")

	if err != nil {
		panic(fmt.Errorf("Unable find config directory: %s \n", err))
	}

	viper.SetDefault("env", "development")

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFile)
	env := viper.GetString("env")
	viper.SetConfigName(env)

	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}
