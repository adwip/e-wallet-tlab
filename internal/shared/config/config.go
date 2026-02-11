package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	Service  ServiceConfig  `mapstructure:"service" yaml:"service"`
	Db       DbConfig       `mapstructure:"db" yaml:"db"`
	Security SecurityConfig `mapstructure:"security" yaml:"security"`
}

type ServiceConfig struct {
	Name    string `mapstructure:"name" yaml:"name"`
	Port    string `mapstructure:"port" yaml:"port"`
	LogFile string `mapstructure:"log_file" yaml:"log_file"`
}

type DbConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
}

type SecurityConfig struct {
	PasswordHashKey string `mapstructure:"password_hash_key" yaml:"password_hash_key"`
}

func SetupConfig() (out ConfigApp, err error) {
	err = readLocalConfig()
	if err != nil {
		return out, err
	}

	viper.Unmarshal(&out)
	return out, nil
}

func readLocalConfig() (err error) {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("Path root:", path)
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(path, "internal/shared/config"))
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
