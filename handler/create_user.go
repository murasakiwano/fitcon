package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/fitconner"
	"go.uber.org/zap"
)

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(fitconner.FitConner)
	if err := c.Bind(u); err != nil {
		return err
	}

	h.log.Debugw("Received", zap.Any("user", u))

	if err := h.db.CreateFitConner(*u); err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"matricula": u.ID,
		"name":      u.Name,
	})
}
