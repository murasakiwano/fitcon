package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/components"
	"github.com/murasakiwano/fitcon/db"
	"github.com/murasakiwano/fitcon/fitconner"
	"go.uber.org/zap"
)

type Handler struct {
	db  *db.DB
	log *zap.SugaredLogger
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

	r := struct {
		Id   string `json:"matricula"`
		Name string `json:"name"`
	}{
		Id:   u.ID,
		Name: u.Name,
	}

	return c.JSON(http.StatusCreated, r)
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
	if err := components.Index(components.Home()).Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}

func (h *Handler) GetHome(c echo.Context) error {
	if err := components.Home().Render(c.Request().Context(), c.Response().Writer); err != nil {
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
