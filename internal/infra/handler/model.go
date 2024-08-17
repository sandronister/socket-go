package handler

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
)

const bufferMaxSize = 1024

type TcpHandler struct {
	broker     types.IBroker
	ReadBuffer []byte
	Retries    int
	MaxRetries int
}
