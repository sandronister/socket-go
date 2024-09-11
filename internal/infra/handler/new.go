package handler

import "github.com/sandronister/socket-go/pkg/devices/usecase"

func NewTcpHandler(device usecase.IDeviceUseCase, maxRetries int) IHandler {
	return &TcpHandler{
		udevice:    device,
		MaxRetries: maxRetries,
		Retries:    0,
	}
}
