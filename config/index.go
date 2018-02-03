package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Port string
	}
}

var C config

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
	env := viper.Get("env").(string)
	viper.SetConfigName(env)

	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Println("Unable to decode into struct, %v", err)
	}
}
