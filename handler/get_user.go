package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/components"
	"go.uber.org/zap"
)

func (h *Handler) GetUser(c echo.Context) error {
	h.log.Debugw("Request",
		zap.Any("path", c.Request().URL.Path),
		zap.Any("params", c.QueryParams()),
	)
	id := c.QueryParam("matricula")
	h.log.Debugw("Got", zap.String("matricula", id))

	if err := h.db.ValidateId(id); err != nil {
		h.log.Error(err)
		return components.UserIdInvalid(id).Render(
			context.Background(),
			c.Response().Writer,
		)
	}

	fc, err := h.db.GetFitConner(id)
	if err != nil {
		h.log.Error(err)
		return components.UserNotFound(id).Render(
			c.Request().Context(),
			c.Response().Writer,
		)
	}

	if err := components.UserTable(*fc).Render(
		c.Request().Context(),
		c.Response().Writer,
	); err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}
