package ruptela

import (
	"bytes"

	"github.com/sandronister/socket-go/pkg/devices/entities/abstract"
)

func New(buffer []byte) *Device {
	return &Device{
		Device: abstract.Device{
			OriginalBuff: buffer,
			Buffer:       bytes.NewBuffer(buffer),
			BytesReaded:  0,
		},
	}
}
