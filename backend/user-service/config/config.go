package config

import (
	"github.com/spf13/viper"
)

type UserServiceConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	UserService UserServiceConfig `yaml:"user_service"`
	Database    DatabaseConfig    `yaml:"database"`
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
