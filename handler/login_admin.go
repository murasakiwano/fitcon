package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo-contrib/session"
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
			"Errorrrr",
			zap.String("error", "admin secret does not match"),
			zap.String("got", params.AdminSecret),
			zap.String("expected", os.Getenv("ADMIN_SECRET")),
		)
		return echo.ErrUnauthorized
	}

	h.log.Debugw("params",
		zap.String("name", params.Name),
		zap.String("password", params.Password),
	)

	sess, _ := session.Get(params.Name, c)
	if sess.Values["authenticated"] == true {
		// TODO: REDIRECT
		return nil
	}

	token, err := auth.MakeJWT(params.Name, h.jwtSecret, time.Duration(60*60)*time.Second, "fitcon", true)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	sess.Values["admin"] = true
	sess.Values["token"] = token
	sess.Values["authenticated"] = true
	h.log.Infow("got session",
		zap.String("id", sess.ID),
		zap.Any("options", sess.Options),
		zap.Bool("isNew", sess.IsNew),
	)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"name":  params.Name,
	})
}
