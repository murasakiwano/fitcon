package handler

import (
	// "fmt"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/auth"
	"github.com/murasakiwano/fitcon/internal/db"
	"go.uber.org/zap"
	// "golang.org/x/crypto/bcrypt"
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

// Creates a user with admin roles
func (h *Handler) CreateAdmin(c echo.Context) error {
	var params adminParameters
	if err := c.Bind(&params); err != nil {
		h.log.Error(zap.Error(err))
		return echo.ErrBadRequest
	}
	if params.AdminSecret != os.Getenv("ADMIN_SECRET") {
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

	token, err := auth.MakeJWT(params.Name, h.jwtSecret, time.Duration(60*60)*time.Second, "fitcon", true)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	sess, _ := session.Get(params.Name, c)
	sess.Options = &DefaultOptions

	sess.Values["token"] = token
	sess.Values["authenticated"] = true

	return c.JSON(http.StatusCreated, echo.Map{
		"token":   token,
		"id":      params.Name,
		"message": "successfully created user",
	})
}

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
