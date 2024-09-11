package sugar

import (
	"github.com/sandronister/socket-go/pkg/logger/types"
)

type Logger struct {
	sugar types.ISuggarEngine
}

func NewLogger(sugar types.ISuggarEngine) types.ILogger {
	return &Logger{sugar: sugar}
}
