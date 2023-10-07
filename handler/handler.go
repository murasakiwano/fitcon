package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/components"
	"github.com/murasakiwano/fitcon/db"
	"github.com/murasakiwano/fitcon/fitconner"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type Handler struct {
	FitConnerStore *db.FitConnerStore
	Log            *slog.Logger
}

func New(log *slog.Logger, fcs *db.FitConnerStore) *Handler {
	return &Handler{
		Log:            log,
		FitConnerStore: fcs,
	}
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.QueryParam("matricula")
	c.Logger().Debug("matricula = " + id)

	h.FitConnerStore.GetFitconner(id)

	if err := components.UserTable("Zezin da Viola").Render(c.Request().Context(), c.Response().Writer); err != nil {
		c.Logger().Error(err)
		return err
	}

	return nil
}

func (h *Handler) GetIndex(c echo.Context) error {
	if err := components.Index().Render(c.Request().Context(), c.Response().Writer); err != nil {
		c.Logger().Error(err)
		return err
	}

	return nil
}

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(fitconner.FitConner)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := h.FitConnerStore.InsertFitconner(*u); err != nil {
		h.Log.Error("%v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, u)
}
