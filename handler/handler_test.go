package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/murasakiwano/fitcon/db"
	"github.com/murasakiwano/fitcon/fitconner"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	log, _   = zap.NewDevelopment()
	sugar    = log.Sugar()
	userJSON = `{"matricula":"C123456","name":"John Doe"}
`
	fc = fitconner.New(
		"C123456",
		"John Doe",
		"Team 1",
		"10",
		"20",
		"30",
		"40",
		"50",
		1,
	)
)

var sldb *sqlx.DB

func init() {
	var err error
	sldb, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("Error connecting to the database", err)
	}
}

var h Handler

func init() {
	fcs, err := db.New(sugar, ":memory:")
	if err != nil {
		sugar.Error("failed to create store", zap.Error(err))
	}

	h = New(sugar, fcs)
}

func TestGetFitconner(t *testing.T) {
	// Setup
	h.db.Create()
	defer h.db.Drop()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users?matricula=C123456", nil)
	rec := httptest.NewRecorder()
	err := h.db.CreateFitConner(*fc)
	if err != nil {
		slog.Error("Error:", err)
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
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("matricula", "name", "teamName", "teamNumber", "goal1", "goal2")
	c.SetParamValues(
		fc.ID,
		strconv.Itoa(fc.TeamNumber),
		fc.TeamName,
		fc.Goal1FatPercentage,
		fc.Goal1LeanMass,
		fc.Goal2VisceralFat,
		fc.Goal2FatPercentage,
		fc.Goal2LeanMass,
	)

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
