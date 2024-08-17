package handler

import "github.com/sandronister/go_broker/pkg/broker/types"

func NewTcpHandler(broker types.IBroker, maxRetries int) *TcpHandler {
	return &TcpHandler{
		broker:     broker,
		ReadBuffer: make([]byte, bufferMaxSize),
		MaxRetries: maxRetries,
		Retries:    0,
	}
}
