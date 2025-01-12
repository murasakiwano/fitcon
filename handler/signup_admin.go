package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/auth"
	"github.com/murasakiwano/fitcon/internal/db"
	"go.uber.org/zap"
)

func (h *Handler) CreateAdmin(c echo.Context) error {
	var params adminParameters
	if err := c.Bind(&params); err != nil {
		h.log.Error(zap.Error(err))
		return echo.ErrBadRequest
	}

	if params.AdminSecret != os.Getenv("ADMIN_SECRET") {
		h.log.Debugw("Error comparing secrets",
			zap.String("got", params.AdminSecret),
			zap.String("expected", os.Getenv("ADMIN_SECRET")),
		)
		return echo.ErrUnauthorized
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		return err
	}

	err = h.db.CreateAdmin(db.Admin{
		Name:           params.Name,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		h.log.Errorw("error creating admin", zap.Error(err))
		if errors.Is(err, db.ErrAlreadyExists) {
			return c.JSON(http.StatusConflict, "admin already exists")
		}

		return c.String(http.StatusInternalServerError, "Couldn't create admin")
	}

	sess, _ := session.Get(SessionName, c)
	sess.Options = &DefaultOptions

	sess.Values["admin"] = true
	sess.Values["user_name"] = params.Name
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response().Writer)

	return h.GetIndex(c)
}
