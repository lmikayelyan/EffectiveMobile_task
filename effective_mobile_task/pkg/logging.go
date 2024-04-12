package pkg

import (
	"context"
	"effective_mobile_task/internal/config"
	"github.com/rs/zerolog"
	"log"
	"os"
)

type Logger interface {
	InitLogger(ctx context.Context) (zerolog.Logger, error)
}

type logger struct {
	cfg *config.Config
}

func NewLogger(cfg *config.Config) Logger {
	return &logger{cfg: cfg}
}

func (l *logger) InitLogger(ctx context.Context) (zerolog.Logger, error) {
	newLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logLevel, err := zerolog.ParseLevel(l.cfg.LogLevel)
	if err != nil {
		log.Panic(err)
	}

	newLogger.Level(logLevel)

	return newLogger, nil
}
