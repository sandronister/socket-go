package handler

import "github.com/sandronister/socket-go/pkg/devices/usecase"

func NewTcpHandler(device usecase.IDeviceUseCase, maxRetries int) *TcpHandler {
	return &TcpHandler{
		udevice:    device,
		MaxRetries: maxRetries,
		Retries:    0,
	}
}
