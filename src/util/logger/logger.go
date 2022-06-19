package logger

import (
	"github.com/betalixt/eventSourceUsers/util/blerr"
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {	
		panic(blerr.NewError(blerr.LoggerCreateFailureCode, 500, err.Error()))
	}
	return logger
}
