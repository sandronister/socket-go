package ruptela

import "github.com/sandronister/socket-go/pkg/devices/entities/abstract"

type KeyCountRuptela struct {
	Data  string
	Count int
}

type Header struct {
	Imei string
}

type Device struct {
	abstract.Device
	Header  Header
	Ack     []byte
	Success bool
}
