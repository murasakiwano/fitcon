package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *Handler) UpdateUser(c echo.Context) error {
	params, err := c.FormParams()
	if err != nil {
		h.log.Error(zap.Error(err))
		return err
	}

	id := params.Get("matricula")
	if id == "" {
		h.log.Error("No matricula provided")
		return c.JSON(http.StatusBadRequest, "No matricula provided")
	}

	params.Del("matricula")

	mapParams := make(map[string]string)
	for k, v := range params {
		mapParams[k] = v[0]
	}

	_, err = h.db.UpdateFitConner(id, mapParams)
	if err != nil {
		return err
	}

	jsonParams, err := json.Marshal(mapParams)
	if err != nil {
		h.log.Error("Error marshaling params", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"matricula": id,
		"updated":   jsonParams,
	})
}
