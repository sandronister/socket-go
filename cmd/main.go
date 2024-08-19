package main

import (
	"sync"

	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/di"
	"github.com/sandronister/socket-go/pkg/catch"
)

func main() {
	_, err := config.LoadConfig(".")
	catch.Exception(err)

	catch.Exception(err)

	wg := sync.WaitGroup{}
	wg.Add(1)

	server, err := di.NewServer()
	catch.Exception(err)

	go server.Start()

	wg.Wait()

}
