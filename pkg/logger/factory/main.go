package factory

import (
	"fmt"
	"time"

	"github.com/sandronister/socket-go/pkg/logger/sugar"
	"github.com/sandronister/socket-go/pkg/logger/types"
	"go.uber.org/zap"
)

type LoggerType string

const (
	Sugar LoggerType = "sugar"
)

func newSugar(pattern string) (types.ILogger, error) {
	timePattern := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s/%s.log", pattern, timePattern)

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{fileName}
	logger, err := config.Build()

	if err != nil {
		return nil, err
	}

	objSug := logger.Sugar()

	return sugar.NewLogger(objSug), nil
}

func NewLogger(loggerType LoggerType, pattern string) (types.ILogger, error) {
	switch loggerType {
	case Sugar:
		return newSugar(pattern)
	default:
		return nil, fmt.Errorf("logger type %s not supported", loggerType)
	}
}
