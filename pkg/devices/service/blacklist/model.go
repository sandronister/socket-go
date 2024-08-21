package blacklist

import "github.com/sandronister/socket-go/pkg/devices/repositories"

type blacklistservice struct {
	repository      repositories.IDeviceRepository
	blackListDevice map[string]bool
}
