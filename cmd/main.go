package main

import (
	"github.com/sandronister/go_broker/pkg/kafka"
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/infra/web"
	"github.com/sandronister/socket-go/pkg/catch"
)

func main() {
	conf, err := config.LoadConfig(".")
	catch.Exception(err)

	catch.Exception(err)

	broker := kafka.NewBroker(conf.BROKER_HOST, conf.BROKER_PORT)

	server := web.NewServer(conf, broker)
	server.Start()

}
