package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	h.db.Create()
	defer h.db.Drop()
	fc.Password = ""
	h.db.CreateFitConner(*fc)
	e := echo.New()
	f := make(url.Values)
	f.Set("matricula", fc.ID)
	f.Set("password", fc.Password)
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"id":"C123456"}`+"\n", rec.Body.String())
	}
}

func TestLogsInCorrectly(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	fitConner := fc
	fitConner.Password = "test-password1"
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
