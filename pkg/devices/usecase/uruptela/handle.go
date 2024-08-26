package uruptela

import (
	"bytes"

	"github.com/sandronister/socket-go/internal/dto"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *UseRuptela) Handle(buff []byte) *dto.DeviceResponse {
	r := ruptela.New(buff)
	r.Buffer = bytes.NewBuffer(buff)
	behave, err := u.ProcessBehave(r)
	return &dto.DeviceResponse{
		Error:          err,
		Ack:            r.Ack,
		ProtocolBehave: behave,
		Success:        r.Success,
		Imei:           r.Header.Imei,
	}
}
