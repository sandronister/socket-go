package blacklist

import (
	"github.com/sandronister/socket-go/pkg/devices/repositories"
	"github.com/sandronister/socket-go/pkg/devices/service"
)

var blacklist service.IBlackListService

func Factory(repository repositories.IDeviceRepository) service.IBlackListService {
	if blacklist == nil {
		blacklist = New(repository)
	}
	return blacklist
}
