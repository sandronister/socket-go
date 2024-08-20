package usecase

import "github.com/sandronister/socket-go/internal/dto"

type IUseCase interface {
	Handle(buff []byte) *dto.DeviceResponse
}
