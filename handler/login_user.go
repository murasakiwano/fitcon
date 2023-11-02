package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/components"
	"github.com/murasakiwano/fitcon/internal/auth"
	"go.uber.org/zap"
)

func (h *Handler) Login(c echo.Context) error {
	type User struct {
		ID       string `json:"matricula" form:"matricula"`
		Password string `json:"password" form:"password"`
	}
	type response struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		// User
	}

	var user User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	participant, err := h.db.GetFitConner(user.ID)
	if err != nil {
		return c.String(http.StatusUnauthorized, "user does not exist")
	}

	err = auth.CheckPasswordHash(user.Password, participant.HashedPassword)
	if err != nil {
		return h.renderComponent(
			components.Index(
				components.InfoPage("Senha errada",
					"A senha inserida está incorreta"),
			),
			c,
		)
	}

	sess, err := h.createSession(user.ID, c)
	if err != nil {
		h.log.Errorw("error occurred when creating session", zap.Error(err))
		return echo.ErrInternalServerError
	}

	if sess.Values["authenticated"] == true {
		return h.GetIndex(c)
	}

	sess.Values["user_id"] = user.ID
	sess.Values["admin"] = false
	sess.Values["authenticated"] = true
	// Sends a SetCookie header back with te cookie value being the session name
	sess.Save(c.Request(), c.Response())

	return h.GetIndex(c)
}
