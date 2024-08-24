package main

import (
	"os"
	"sync"

	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/di"
	"github.com/sandronister/socket-go/pkg/catch"
	"gitlab.com/gobrax-dev/gobrax-tool/cloud/factory"
)

func main() {
	err := config.LoadConfig(".env")
	catch.Exception(err)

	catch.Exception(err)

	wg := sync.WaitGroup{}
	wg.Add(1)

	server, err := di.NewServer()
	catch.Exception(err)

	go server.Start()

	cloud, err := factory.GetCloudInstance()

	if err != nil {
		panic(err)
	}

	bucket := cloud.GetBucket()

	list, err := bucket.ListObject(os.Getenv("BUCKET_NAME"))

	if err != nil {
		panic(err)
	}

	for _, obj := range list {
		println(obj)
	}

	wg.Wait()

}
