package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/components"
)

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
