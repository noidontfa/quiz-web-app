package config

import (
	"github.com/spf13/viper"
)

type config struct {
}

type DatabaseConfigurations struct {
	DBName         string
	DBUser         string
	DBPassword     string
	Port           string
	HttpServerHost string
}

func NewConfiguration() *config {
	return &config{}
}

func (c *config) GetConfig() (*DatabaseConfigurations, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./src")
	viper.SetConfigType("yml")
	var configuration DatabaseConfigurations

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
