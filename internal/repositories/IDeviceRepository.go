package repositories

import "github.com/sandronister/socket-go/internal/entities"

type IDeviceRepository interface {
	SaveMessage(md5Str string)
	GetBlackList() ([]entities.BlackListDevice, error)
}
