package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	h.db.Create()
	defer h.db.Drop()
	fc.ClearPassword()
	h.db.CreateFitConner(*fc)
	e := echo.New()
	f := make(url.Values)
	store := sessions.NewCookieStore([]byte("secret"))
	e.Use(session.Middleware(store))
	f.Set("matricula", fc.ID)
	f.Set("password", fc.HashedPassword)
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"id":"C123456"}`+"\n", rec.Body.String())
	}
}

func TestLogin(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	fitConner := fc
	fitConner.SetPassword("test-password1")
	h.db.CreateFitConner(*fitConner)
	e := echo.New()
	f := make(url.Values)
	f.Set("username", fc.ID)
	f.Set("password", "test-password1")
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestMultipleSessions(t *testing.T) {
}
