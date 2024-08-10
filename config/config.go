package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Conf struct {
	Device map[string]string
}

func LoadConfig(path string) (*Conf, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	mapEnv := viper.GetString("DEVICES")
	mapData := make(map[string]string)
	for _, v := range strings.Split(mapEnv, ",") {
		split := strings.Split(v, ":")
		if len(split) == 2 {
			mapData[split[0]] = split[1]
		}
	}

	return &Conf{
		Device: mapData,
	}, nil

}
