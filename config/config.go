package config

import (
	"encoding/json"
	"os"
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

func LoadDevices(path string) (*Conf, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Conf
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
