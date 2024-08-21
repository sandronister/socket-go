package usecase

import "github.com/sandronister/socket-go/internal/dto"

type IDeviceUseCase interface {
	Handle(buff []byte) *dto.DeviceResponse
}
