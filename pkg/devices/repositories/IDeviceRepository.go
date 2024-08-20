package repositories

import "github.com/sandronister/socket-go/pkg/devices/types"

type IDeviceRepository interface {
	GetBlackList() ([]types.BlackListDevice, error)
}
