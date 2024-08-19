package repositories

import "github.com/sandronister/socket-go/internal/entities"

type IDeviceRepository interface {
	SaveMessage(device entities.Ruptela) error
	GetBlackList() ([]entities.BlackListDevice, error)
}
