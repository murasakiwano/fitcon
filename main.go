package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/murasakiwano/fitcon/db"
	"github.com/murasakiwano/fitcon/handler"
	"go.uber.org/zap"
)

func main() {
	var logger zap.SugaredLogger

	if os.Getenv("ENV") == "production" {
		l, _ := zap.NewProduction()
		logger = *l.Sugar()
	} else {
		l, _ := zap.NewDevelopment()
		logger = *l.Sugar()
	}

	fcs, err := db.New(&logger, db.FitConnersTable)
	defer fcs.CloseDB()
	if err != nil {
		logger.Error(zap.Error(err))
		os.Exit(1)
	}

	h := handler.New(&logger, fcs)

	Serve(&h)
}

func Serve(h *handler.Handler) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/assets", "assets")
	e.Static("/img", "img")
	e.Static("/css", "css")
	e.File("/favicon.ico", "favicon.ico")

	e.GET("/", h.GetIndex)
	e.GET("/users", h.GetUser)
	e.POST("/users", h.CreateUser)
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":1323"))
}
