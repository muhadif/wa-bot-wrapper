package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Username string `mapstructure:"DATABASE_USERNAME"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     string `mapstructure:"DATABASE_PORT"`
	Database string `mapstructure:"DATABASE_DATABASE"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	// Automatically read in environment variables that match
	viper.AutomaticEnv()

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Create an instance of the Config struct
	var config Config
	err := viper.Unmarshal(&config, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	)))
	if err != nil {
		return nil, err
	}

	return &config, nil
}
