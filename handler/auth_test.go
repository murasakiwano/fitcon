package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/auth"
	"github.com/murasakiwano/fitcon/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	h.db.Create()
	defer h.db.Drop()

	fc.ClearPassword()
	h.db.CreateFitConner(*fc)

	e := echo.New()
	store := sessions.NewCookieStore([]byte("secret"))
	mw := session.Middleware(store)
	handler := mw(h.SignUp)

	f := make(url.Values)
	f.Set("matricula", fc.ID)
	f.Set("password", fc.HashedPassword)

	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"id":"C123456"}`+"\n", rec.Body.String())
		assertSession(t, fc.ID, c)
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

	mw := session.Middleware(store)
	handler := mw(h.Login)

	f := make(url.Values)
	f.Set("matricula", fc.ID)
	f.Set("password", "test-password1")
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assertSession(t, fc.ID, c)
	}
}

func TestAdminSignUp(t *testing.T) {
	h.db.Create()
	defer h.db.Drop()

	name := "Zequinha"
	password := "admin_password!00"
	_, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("Error creating hash: %s", err)
	}

	e := echo.New()
	store := sessions.NewCookieStore([]byte("secret"))
	mw := session.Middleware(store)
	hand := mw(h.CreateAdmin)

	f := make(url.Values)
	f.Set("username", name)
	f.Set("password", password)
	f.Set("admin_secret", "opa")

	os.Setenv("ADMIN_SECRET", "opa")
	req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, hand(c)) {
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		admin, err := h.db.GetAdmin(name)
		assert.Nil(t, err)
		assert.Equal(t, admin.Name, name)

		assert.NoError(t, auth.CheckPasswordHash(password, admin.HashedPassword))
		assertSession(t, name, c)
	}
}

func TestAdminLogin(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()

	name := "Zequinha"
	password := "admin_password!00"
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("Error creating hash: %s", err)
	}

	h.db.CreateAdmin(db.Admin{
		Name:           name,
		HashedPassword: hashedPassword,
	})

	e := echo.New()
	store := sessions.NewCookieStore([]byte("secret"))
	mw := session.Middleware(store)
	handler := mw(h.LoginAdmin)

	f := make(url.Values)
	f.Set("username", name)
	f.Set("password", password)
	f.Set("admin_secret", "opa")

	os.Setenv("ADMIN_SECRET", "opa")
	req := httptest.NewRequest(http.MethodPost, "/admin/login", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assertSession(t, name, c)
	}
}

func assertSession(t *testing.T, sessionName string, c echo.Context) {
	sess, err := session.Get(SessionName, c)
	assert.Nil(t, err)
	assert.True(t, sess.Values["authenticated"].(bool))
	assert.True(t, sess.Options.HttpOnly)
}
