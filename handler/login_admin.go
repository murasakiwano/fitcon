package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/auth"
	"go.uber.org/zap"
)

func (h *Handler) LoginAdmin(c echo.Context) error {
	var params adminParameters
	if err := c.Bind(&params); err != nil {
		h.log.Error(zap.Error(err))
		return echo.ErrBadRequest
	}
	if params.AdminSecret != os.Getenv("ADMIN_SECRET") {
		h.log.Errorw(
			"Error encountered",
			zap.String("error", "admin secret does not match"),
		)
		return echo.ErrUnauthorized
	}

	sess, err := h.createSession(params.Name, c)
	if err != nil {
		h.log.Errorw("error occurred when creating session", zap.Error(err))
		return echo.ErrInternalServerError
	}

	token, err := auth.MakeJWT(params.Name, h.jwtSecret, time.Duration(60*60)*time.Second, "fitcon", true)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	sess.Values["admin"] = true
	sess.Values["token"] = token
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response().Writer)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"name":  params.Name,
	})
}
