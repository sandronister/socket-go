package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	viper.AutomaticEnv()
	return nil
}
