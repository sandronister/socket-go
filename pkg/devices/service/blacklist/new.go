package blacklist

import (
	"github.com/sandronister/socket-go/pkg/devices/repositories"
	"github.com/sandronister/socket-go/pkg/devices/service"
)

func New(repository repositories.IDeviceRepository) service.IBlackListService {
	blacklist := &blacklistservice{
		repository: repository,
	}
	blacklist.InitBlackListDevice()
	return blacklist
}
