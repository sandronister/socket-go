package main

import (
	"sync"

	brokerredis "github.com/sandronister/go_broker/pkg/broker_redis"
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/infra/web"
	"github.com/sandronister/socket-go/pkg/catch"
)

func main() {
	conf, err := config.LoadConfig(".")
	catch.Exception(err)

	wg := sync.WaitGroup{}
	wg.Add(1)

	broker := brokerredis.NewBroker(conf.BROKER_HOST, conf.BROKER_PORT)

	server := web.NewServer(conf, broker)
	go server.Start()

	wg.Wait()

}
