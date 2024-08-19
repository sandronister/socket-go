package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	BROKER_HOST  string `mapstructure:"BROKER_HOST"`
	BROKER_PORT  int    `mapstructure:"BROKER_PORT"`
	SOCKET_HOST  string `mapstructure:"SOCKET_HOST"`
	SOCKET_PORT  string `mapstructure:"SOCKET_PORT"`
	DEVICE_TOPIC string `mapstructure:"DEVICE_TOPIC"`
	TIME_FLUSH   int    `mapstructure:"TIME_FLUSH"`
	DatabaseURL  string `mapstructure:"DatabaseURL"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("kafka")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
