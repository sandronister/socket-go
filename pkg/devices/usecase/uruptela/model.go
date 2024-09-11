package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices/usecase/abstract"
	tcloud "gitlab.com/gobrax-dev/gobrax-tool/cloud/types"
)

type UseRuptela struct {
	abstract.AbstractUseCase
	cloud tcloud.ICloudProvider
}
