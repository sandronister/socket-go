package main

import (
	"sync"

	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/di"
	"github.com/sandronister/socket-go/pkg/catch"
)

func main() {
	conf, err := config.LoadConfig(".")
	catch.Exception(err)

	wg := sync.WaitGroup{}
	wg.Add(1)

	server := di.NewServer(conf)
	go server.Start()

	wg.Wait()

}
