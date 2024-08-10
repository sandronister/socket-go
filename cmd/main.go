package main

import (
	"fmt"

	"github.com/sandronister/socket-go/config"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
		return

	}
	for k, v := range config.Device {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	// server := web.NewServer("localhost", "8080")
	// server.Start()
}
