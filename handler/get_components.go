package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/components"
)

func (h *Handler) GetIndex(c echo.Context) error {
	comp := components.Index(components.GetUser())
	return h.renderComponent(comp, c)
}

func (h *Handler) GetCreate(c echo.Context) error {
	comp := components.Index(components.CreateUser())
	return h.renderComponent(comp, c)
}

func (h *Handler) GetSignUp(c echo.Context) error {
	comp := components.Index(components.SignUp())
	return h.renderComponent(comp, c)
}

func (h *Handler) GetLogin(c echo.Context) error {
	comp := components.Index(components.Login())
	return h.renderComponent(comp, c)
}

func (h *Handler) GetLoginAdmin(c echo.Context) error {
	comp := components.Index(components.LoginAdmin())
	return h.renderComponent(comp, c)
}

func (h *Handler) GetSignUpAdmin(c echo.Context) error {
	comp := components.Index(components.SignUpAdmin())
	return h.renderComponent(comp, c)
}

func (h *Handler) renderComponent(comp templ.Component, c echo.Context) error {
	if err := comp.Render(c.Request().Context(), c.Response().Writer); err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}
