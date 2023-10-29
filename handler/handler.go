package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/components"
	"github.com/murasakiwano/fitcon/internal/db"
	"github.com/murasakiwano/fitcon/internal/fitconner"
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
		"nome":      u.Name,
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	h.log.Debugw("Request", zap.Any("path", c.Request().URL.Path), zap.Any("context", c.Request().Context()), zap.Any("body", c.Request().Body))
	id := c.QueryParam("matricula")
	h.log.Debugw("Got", zap.String("matricula", id))

	if err := h.db.ValidateId(id); err != nil {
		h.log.Error(err)
		return components.UserIdInvalid(id).Render(context.Background(), c.Response().Writer)
	}

	fc, err := h.db.GetFitConner(id)
	if err != nil {
		h.log.Error(err)
		return components.UserNotFound(id).Render(c.Request().Context(), c.Response().Writer)
	}

	if err := components.UserTable(*fc).Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

func (h *Handler) GetIndex(c echo.Context) error {
	if err := components.Index(components.Home(components.Login())).Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}

func (h *Handler) GetHome(c echo.Context) error {
	if err := components.Home(components.GetUserForm()).Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}

func (h *Handler) GetCreate(c echo.Context) error {
	if err := components.Index(components.CreateUser()).Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}

func (h *Handler) DeleteUser(c echo.Context) error {
	h.log.Debugw("Request", zap.Any("path", c.Request().URL.Path), zap.Any("context", c.Request().Context()), zap.Any("body", c.Request().Body))
	id := c.QueryParam("matricula")
	h.log.Debugw("Got", zap.String("matricula", id))

	err := h.db.DeleteFitConner(id)
	if err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := c.JSON(http.StatusOK, fmt.Sprintf(`{"deleted":"%s"}`, id))
	h.log.Debugw("Response", zap.Any("response", res))
	return res
}

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

	return c.JSON(http.StatusOK, fmt.Sprintf(`{"matricula":"%s","updated":%s}`, id, jsonParams))
}
