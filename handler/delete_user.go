package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *Handler) DeleteUser(c echo.Context) error {
	h.log.Debugw("Request", zap.Any("path", c.Request().URL.Path), zap.Any("context", c.Request().Context()), zap.Any("body", c.Request().Body))
	id := c.QueryParam("matricula")
	h.log.Debugw("Got", zap.String("matricula", id))

	err := h.db.DeleteFitConner(id)
	if err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := c.JSON(http.StatusOK, echo.Map{
		"deleted": id,
	})
	h.log.Debugw("Response", zap.Any("response", res))
	return res
}
