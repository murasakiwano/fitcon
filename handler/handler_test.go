package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/internal/db"
	"github.com/murasakiwano/fitcon/internal/fitconner"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

var (
	userJSON = `{"matricula":"C123456","name":"John Doe"}
`
	fc, _ = fitconner.New(
		"C123456",
		"John Doe",
		"",
		"Team 1",
		"10",
		"20",
		"30",
		"40",
		"50",
		"10",
		"10",
		1,
	)
)

var (
	sldb *sqlx.DB
	h    Handler
)

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := cfg.Build()
	sugar = logger.Sugar()
	os.Setenv("DATABASE_FILE", ":memory:")
	fcs, err := db.New(sugar)
	if err != nil {
		sugar.Error("failed to create store", zap.Error(err))
	}

	h = New(sugar, fcs)
}

func TestGetFitConner(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users?matricula=C123456", nil)
	rec := httptest.NewRecorder()
	err := h.db.CreateFitConner(*fc)
	if err != nil {
		sugar.Error("Error:", err)
	}
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateFitConner(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	e := echo.New()
	f := make(url.Values)
	f.Set("matricula", fc.ID)
	f.Set("name", fc.Name)
	f.Set("teamName", fc.TeamName)
	f.Set("teamNumber", strconv.Itoa(fc.TeamNumber))
	f.Set("goal1FatPercentage", fc.Goal1FatPercentage)
	f.Set("goal1LeanMass", fc.Goal1LeanMass)
	f.Set("goal2VisceralFat", fc.Goal2VisceralFat)
	f.Set("goal2FatPercentage", fc.Goal2FatPercentage)
	f.Set("goal2LeanMass", fc.Goal2LeanMass)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestUpdateFitConner(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	h.db.CreateFitConner(*fc)
	e := echo.New()
	f := make(url.Values)
	f.Set("matricula", fc.ID)
	f.Set("goal2_lean_mass", "Aumentar 2kg")
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := "{\"matricula\":\"C123456\",\"updated\":\"{\\\"goal2_lean_mass\\\":\\\"Aumentar 2kg\\\"}\"}\n"

	// Assertions
	if assert.NoError(t, h.UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestDeleteFitConner(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	h.db.GetAllFitConners()
	err := h.db.CreateFitConner(*fc)
	if err != nil {
		t.Fatalf("Error creating fitconner: %s", err)
	}
	e := echo.New()
	q := make(url.Values)
	q.Set("matricula", fc.ID)
	req := httptest.NewRequest(http.MethodPut, "/users?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := "{\"deleted\":\"C123456\"}\n"

	// Assertions
	if assert.NoError(t, h.DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
