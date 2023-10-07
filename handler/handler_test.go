package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"name":"Monkey D. Luffy","matricula":"C0123456789"}`
	mockDB   = map[string]*FitConner{
		"C0123456789": NewFitConner("C0123456789", "Monkey D. Luffy", "Mugiwara", "Diminuir 5%", "Manter", "Aumentar 3%", "Aumentar 2kg", "Diminuir 10kg", 1),
	}
)

func TestGetFitconner(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users?matricula=C0123456789")

	if assert.NoError(t, getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestCreateFitConner(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
