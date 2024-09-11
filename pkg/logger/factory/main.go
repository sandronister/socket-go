package factory

import (
	"fmt"
	"os"
	"time"

	"github.com/sandronister/socket-go/pkg/logger/sugar"
	"github.com/sandronister/socket-go/pkg/logger/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerType string

const (
	Sugar LoggerType = "sugar"
)

func createIOFile(pattern string) (string, error) {
	timePattern := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("logs/%s-%s.log", pattern, timePattern)

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			return "", err
		}
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if _, err := os.Create(fileName); err != nil {
			return "", err
		}
	}

	return fileName, nil
}

func newSugar(pattern string) (types.ILogger, error) {
	fileName, err := createIOFile(pattern)

	if err != nil {
		return nil, err
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{fileName}
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build(zap.AddCallerSkip(1))

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
