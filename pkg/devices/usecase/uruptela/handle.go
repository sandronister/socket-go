package uruptela

import (
	"github.com/sandronister/socket-go/internal/dto"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *usecase_t) Handle(buff []byte) *dto.DeviceResponse {
	r := ruptela.New(buff)
	return &dto.DeviceResponse{
		Error:          "",
		Ack:            r.Ack,
		ProtocolBehave: 0,
		Success:        r.Success,
		Imei:           r.Header.Imei,
	}
}
