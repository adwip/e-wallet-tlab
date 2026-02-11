package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	Service ServiceConfig `mapstructure:"service" yaml:"service"`
	Db      DbConfig      `mapstructure:"db" yaml:"db"`
}

type ServiceConfig struct {
	Name    string `mapstructure:"name" yaml:"name"`
	Port    string `mapstructure:"port" yaml:"port"`
	LogFile string `mapstructure:"log_file" yaml:"log_file"`
}

type DbConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
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

	// Try to read config.json first
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(path, "internal/shared/config"))
	err = viper.ReadInConfig()

	// If config.json not found, try to load from .env file
	if err != nil {
		fmt.Println("config.json not found, trying to load from .env file...")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(path)

		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println("Failed to read .env file, trying environment variables...")
			// If .env file also not found, use environment variables
			viper.AutomaticEnv()

			// Set default values if needed
			setDefaultConfig()
			return nil
		}
		fmt.Println("Loaded configuration from .env file")
	} else {
		fmt.Println("Loaded configuration from config.json")
	}

	return nil
}

func setDefaultConfig() {
	// Service defaults
	viper.SetDefault("service.name", "e-wallet-tlab")
	viper.SetDefault("service.port", "8080")
	viper.SetDefault("service.log_file", "logs/app.log")

	// Database defaults
	viper.SetDefault("db.host", "localhost:3306")

	// Security defaults
	viper.SetDefault("security.password_hash_key", "default-secret-key")
}
