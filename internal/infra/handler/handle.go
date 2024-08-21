package handler

import (
	"github.com/sandronister/socket-go/internal/dto"
)

var i int = 1

func (h *TcpHandler) Handle() *dto.DeviceResponse {
	return h.udevice.Handle(h.ReadBuffer)
}
