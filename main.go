package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/murasakiwano/fitcon/db"
	"github.com/murasakiwano/fitcon/handler"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	fcs, err := db.NewFitConnerStore("fitcon_goals")
	if err != nil {
		log.Error("failed to create store", slog.Any("error", err))
		os.Exit(1)
	}

	h := handler.New(log, fcs)

	Serve(h)
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
