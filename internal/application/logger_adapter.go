package application

import (
	"log/slog"

	"github.com/leninner/shared/logger"
)

func CreateLoggerAdapter(slogLogger *slog.Logger) *logger.Logger {
	config := logger.LoggerConfig{
		Level:       "info",
		Environment: "development",
		ServiceName: "order-service",
		Encoding:    "console",
	}
	
	loggerInstance, err := logger.NewLogger(config)
	if err != nil {
		panic(err)
	}
	
	return loggerInstance
} 