package handler

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	jwt.RegisteredClaims
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

type adminParameters struct {
	AdminSecret string `json:"admin_secret" form:"admin_secret"`
	Name        string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
}

const (
	DefaultMaxAge = 86400 * 7
)

var DefaultOptions = sessions.Options{
	Path:     "/",
	MaxAge:   86400 * 7,
	HttpOnly: true,
}

var secret = os.Getenv("JWT_SECRET")

func (h *Handler) Logout(c echo.Context) error {
	session, _ := session.Get("fit-session", c)

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(c.Request(), c.Response().Writer)

	return c.JSON(http.StatusOK, "usuário finalizou a sessão")
}
