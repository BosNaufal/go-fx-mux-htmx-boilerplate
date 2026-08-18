package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	DatabaseName     string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabasePort     string

	OpenrouterApiKey string
	Port             string
}

func ReadConfigYaml() *EnvConfig {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := EnvConfig{
		DatabaseName:     viper.GetString("DATABASE_NAME"),
		DatabaseHost:     viper.GetString("DATABASE_HOST"),
		DatabaseUser:     viper.GetString("DATABASE_USER"),
		DatabasePassword: viper.GetString("DATABASE_PASSWORD"),
		DatabasePort:     viper.GetString("DATABASE_PORT"),
		OpenrouterApiKey: viper.GetString("OPENROUTER_API_KEY"),
		Port:             viper.GetString("PORT"),
	}
	return &config
}
