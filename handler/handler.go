package handler

import (
	"github.com/murasakiwano/fitcon/internal/db"
	"go.uber.org/zap"
)

type Handler struct {
	db        *db.DB
	log       *zap.SugaredLogger
	jwtSecret string
}

func New(log *zap.SugaredLogger, fcs *db.DB) Handler {
	return Handler{
		log: log,
		db:  fcs,
	}
}

const SessionName = "fitcon_session"
