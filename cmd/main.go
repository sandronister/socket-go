package main

import "github.com/sandronister/socker-go/internal/infra/web"

func main() {
	server := web.NewServer("localhost", "8080")
	server.Start()
}
