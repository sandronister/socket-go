package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type Device struct {
	Port  string `json:"port"`
	Host  string `json:"host"`
	Topic string `json:"topic"`
	Name  string `json:"name"`
}

type Conf struct {
	Devices []Device `json:"devices"`
}

func LoadConfig() (*Conf, error) {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file")
		return nil, err
	}

	devicesStr := viper.GetString("DEVICES")
	var devices []Device
	err = json.Unmarshal([]byte(devicesStr), &devices)
	if err != nil {
		return nil, err
	}

	return &Conf{
		Devices: devices,
	}, nil
}
