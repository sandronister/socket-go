package main

import (
	"fmt"

	"github.com/sandronister/socket-go/config"
)

func main() {
	config, err := config.LoadDevices("devices.json")
	if err != nil {
		fmt.Println(err)
		return

	}
	for _, device := range config.Devices {
		fmt.Println(device)
	}
	// server := web.NewServer("localhost", "8080")
	// server.Start()
}
