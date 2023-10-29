package handler

import (
	// "fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/auth"
	"go.uber.org/zap"
	// "golang.org/x/crypto/bcrypt"
)

const (
	DefaultMaxAge = 86400 * 7
)

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

	accessToken, err := auth.MakeJWT(user.ID, h.jwtSecret, time.Duration(60*60)*time.Second, "fitcon")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	refreshToken, err := auth.MakeJWT(user.ID, h.jwtSecret, time.Duration(60*60*24*60)*time.Second, "fitcon-refresh")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	// Sends a SetCookie header back with te cookie value being the session name
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, echo.Map{
		"matricula":     user.ID,
		"token":         accessToken,
		"refresh_token": refreshToken,
	})
}

func helper() *echo.Echo {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["foo"] = "bar"
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})

	return e
}

// func (h *Handler) Login(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")
// 	sess, err := session.Get(username, c)
// 	if err != nil {
// 		h.log.Errorw("Error getting session", zap.Error(err))
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"error":   err,
// 			"message": "Erro ao fazer login. Tente novamente mais tarde.",
// 		})
// 	}
// 	if sess == nil {
// 		sess, _ = session.DefaultConfig.Store.New(c.Request(), username)
// 	}
// 	sess.Options = &sessions.Options{
// 		Path:     "/",
// 		MaxAge:   86400 * 7,
// 		HttpOnly: true,
// 	}
//
// 	h.log.Debugw("Logging in", zap.String("username", username))
// 	user, err := h.db.GetFitConner(username)
// 	if err != nil {
// 		h.log.Errorw("Error logging in. User does not exist", zap.Error(err))
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"error":   err,
// 			"message": fmt.Sprintf("Usuário com matrícula %s não existe", username),
// 		})
// 	}
//
// 	if user.PasswordEmpty() {
// 		h.log.Errorw("Error logging in. User does not have a password", zap.Error(err))
// 		return c.JSON(http.StatusUnauthorized, echo.Map{
// 			"message": "User is not registered. Please sign up",
// 		})
// 	}
//
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		h.log.Error("Error: password does not match", zap.Error(err))
// 		return echo.ErrUnauthorized
// 	}
//
// 	sess.Values["userId"] = username
// 	sess.Values["authenticated"] = true
// 	sess.Save(c.Request(), c.Response().Writer)
//
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"logged_in": true,
// 		"username":  username,
// 	})
// }

func (h *Handler) SignUp(c echo.Context) error {
	id := c.FormValue("matricula")
	sess, err := session.Get(id, c)
	if err != nil {
		h.log.Errorw("Error getting session", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err,
			"message": "Erro ao fazer login. Tente novamente mais tarde.",
		})
	}
	if sess == nil {
		sess, _ = session.DefaultConfig.Store.New(c.Request(), id)
	}
	options := sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	h.log.Debugw("Signing up", zap.String("username", id), zap.Any("session_options", options))
	h.log.Debugw("Session", zap.Any("session", sess))
	sess.Options = &options
	password := c.FormValue("password")

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
		return c.JSON(http.StatusInternalServerError, zap.Error(err))
	}

	_, err = h.db.UpdateFitConner(fc.ID, map[string]string{"password": fc.HashedPassword})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, zap.Error(err))
	}

	admin := c.FormValue("adminSecret") == os.Getenv("ADMIN_SECRET")
	sess.Values["admin"] = admin

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
