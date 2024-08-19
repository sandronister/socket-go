package main

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/di"
	"github.com/sandronister/socket-go/pkg/catch"
)

func main() {
	conf, err := config.LoadConfig(".")
	catch.Exception(err)

	conn, err := pgxpool.Connect(context.Background(), conf.DatabaseURL)

	catch.Exception(err)

	wg := sync.WaitGroup{}
	wg.Add(1)

	server := di.NewServer(conf, conn)
	go server.Start()

	wg.Wait()

}
