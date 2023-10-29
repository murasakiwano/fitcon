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
