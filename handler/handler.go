package handler

import (
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

func (h *Handler) GetUser(c echo.Context) error {
	h.log.Debugw("Request", zap.Any("path", c.Request().URL.Path), zap.Any("context", c.Request().Context()), zap.Any("body", c.Request().Body))
	id := c.QueryParam("matricula")
	h.log.Debugw("Got", zap.String("matricula", id))

	fc, err := h.db.GetFitConner(id)
	if err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := components.UserTable(*fc).Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

func (h *Handler) GetIndex(c echo.Context) error {
	if err := components.Index().Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(fitconner.Fitconner)
	if err := c.Bind(u); err != nil {
		return err
	}

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

func (h *Handler) ValidateId(c echo.Context) error {
	id := c.FormValue("matricula")

	if err := h.db.ValidateId(id); err != nil {
		h.log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "ok")
}
