package config

import (
	logger "futurevalue/log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server ServerConfiguration
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.LogErrorF("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.LogErrorF("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
