package handler

import (
	"github.com/sandronister/socket-go/internal/usecase"
)

func NewTcpHandler(usecase usecase.IUseCase, maxRetries int) *TcpHandler {
	return &TcpHandler{
		usecase:    usecase,
		ReadBuffer: make([]byte, bufferMaxSize),
		MaxRetries: maxRetries,
		Retries:    0,
	}
}
