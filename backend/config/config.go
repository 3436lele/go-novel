package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
}

func InitConfig(configPath string) (*Config, error) {
	var config Config

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
