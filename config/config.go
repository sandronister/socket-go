package config

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

type Device struct {
	Port      string `json:"port"`
	Host      string `json:"host"`
	Topic     string `json:"topic"`
	Name      string `json:"name"`
	TimeFlush int    `json:"timeFlush"`
}

type DeviceConf struct {
	Devices []Device `json:"devices"`
}

type Conf struct {
	Broker string `mapstructure:"broker"`
	Port   int    `mapstructure:"port"`
}

func LoadDevices(path string) (*DeviceConf, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config DeviceConf
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
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
