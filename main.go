package main

import (
	"os"

	"github.com/murasakiwano/fitcon/handler"
	"github.com/murasakiwano/fitcon/internal/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	var logger zap.SugaredLogger
	var cfg zap.Config

	if os.Getenv("ENV") == "production" {
		cfg = zap.NewProductionConfig()
		l, _ := cfg.Build()
		logger = *l.Sugar()
	} else {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		l, _ := cfg.Build()
		logger = *l.Sugar()
	}

	fcs, err := db.New(&logger)
	defer fcs.CloseDB()
	if err != nil {
		logger.Error(zap.Error(err))
		os.Exit(1)
	}

	h := handler.New(&logger, fcs)

	h.Serve()
}
