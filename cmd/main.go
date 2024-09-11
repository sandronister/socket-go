package main

import (
	"sync"

	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/di"
	"github.com/sandronister/socket-go/pkg/catch"
	"github.com/sandronister/socket-go/pkg/logger/factory"
)

func main() {
	err := config.LoadConfig(".env")
	catch.Exception(err)

	logger, err := factory.NewLogger(factory.Sugar, "initial")
	catch.Exception(err)

	logger.Info("Starting server")

	wg := sync.WaitGroup{}
	wg.Add(1)

	server, err := di.NewServer()
	catch.Exception(err)

	go server.Start()

	wg.Wait()

}
