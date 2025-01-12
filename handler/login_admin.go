package handler

import (
	"os"

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

	adm, err := h.db.GetAdmin(params.Name)
	if err != nil {
		h.log.Error("Admin does not exist", zap.String("name", params.Name))
		return echo.ErrUnauthorized
	}

	if err := auth.CheckPasswordHash(params.Password, adm.HashedPassword); err != nil {
		h.log.Error("Password does not match")
		return echo.ErrUnauthorized
	}

	sess, err := h.createSession(params.Name, c)
	if err != nil {
		h.log.Errorw("error occurred when creating session", zap.Error(err))
		return echo.ErrInternalServerError
	}

	sess.Values["admin"] = true
	sess.Values["user_name"] = params.Name
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response().Writer)

	return h.GetIndex(c)
}
