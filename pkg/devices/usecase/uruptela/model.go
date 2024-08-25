package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices/usecase/abstract"
	"gitlab.com/gobrax-dev/gobrax-tool/broker/types"
)

type UseRuptela struct {
	abstract.AbstractUseCase
	broker types.IBroker
}
