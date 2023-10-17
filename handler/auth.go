package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/fitconner"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (h *Handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.db.GetFitConner(username)
	if err != nil {
		h.log.Error("Error logging in", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		h.log.Error("Error: password does not match")
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			Issuer:    "fitcon",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	secret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (h *Handler) SignUp(c echo.Context) error {
	id := c.FormValue("matricula")
	password := c.FormValue("password")

	fc, err := h.db.GetFitConner(id)
	if err != nil {
		h.log.Errorw("Error getting fitconner", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, zap.Error(err))
	}

	if fc.Password != "" {
		return c.JSON(http.StatusConflict, "user already has a password")
	}

	fc.Password, err = fitconner.HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, zap.Error(err))
	}

	_, err = h.db.UpdateFitConner(fc.ID, map[string]string{"password": fc.Password})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, zap.Error(err))
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		fc.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			Issuer:    "fitcon",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	secret := os.Getenv("JWT_SECRET")
	_, err = token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id": fc.ID,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
