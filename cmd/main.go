package main

import (
	"sync"

	"github.com/sandronister/go-broker/pkg/kafka"
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/infra/web"
	"github.com/sandronister/socket-go/pkg/catch"
)

func main() {
	conf, err := config.LoadConfig(".")
	catch.Exception(err)

	deviceConf, err := config.LoadDevices("devices.json")
	catch.Exception(err)

	broker := kafka.NewBroker(conf.Broker, conf.Port)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	for _, device := range deviceConf.Devices {
		server := web.NewServer(&device, broker)
		go server.Start()
	}
	waitGroup.Wait()

}
