package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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
		return c.String(http.StatusUnauthorized, "wrong password")
	}

	sess, _ := session.Get(user.ID, c)
	sess.Options = &DefaultOptions
	h.log.Infow("got session",
		zap.String("name", sess.Name()),
		zap.Any("options", sess.Options),
		zap.Bool("isNew", sess.IsNew),
	)

	token, err := auth.MakeJWT(user.ID, h.jwtSecret, time.Duration(60*60)*time.Second, "fitcon", false)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't create access token")
	}
	sess.Values["token"] = token
	sess.Values["authenticated"] = true

	// Sends a SetCookie header back with te cookie value being the session name
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, echo.Map{
		"matricula": user.ID,
		"token":     token,
	})
}