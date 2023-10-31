package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *Handler) SignUp(c echo.Context) error {
	id := c.FormValue("matricula")
	h.log.Debugw("received", zap.String("matricula", id))
	sess, err := session.Get(id, c)
	if err != nil {
		h.log.Errorw("Error getting session", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err,
			"message": "Erro ao fazer login. Tente novamente mais tarde.",
		})
	}

	options := DefaultOptions
	h.log.Debugw("Signing up", zap.String("username", id), zap.Any("session_options", options))
	h.log.Debugw("Session", zap.Any("session", sess))

	sess.Options = &options
	password := c.QueryParam("password")

	fc, err := h.db.GetFitConner(id)
	if err != nil {
		h.log.Errorw("Error getting fitconner", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err,
			"message": "Erro ao cadastrar usuário " + id + ".",
		})
	}

	if !fc.PasswordEmpty() {
		return c.JSON(http.StatusConflict, "usuário já possui senha cadastrada")
	}

	err = fc.SetPassword(password)
	if err != nil {
		h.log.Error(zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	_, err = h.db.UpdateFitConner(fc.ID, map[string]string{"hashed_password": fc.HashedPassword})
	if err != nil {
		h.log.Error(zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response().Writer)

	return c.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
}