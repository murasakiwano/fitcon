package handler

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

var (
	key   = os.Getenv("JWT_SECRET")
	store = sessions.NewCookieStore([]byte(key))
)

func (h *Handler) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))

	e.Static("/assets", "assets")
	e.Static("/img", "img")
	e.Static("/css", "css")

	e.File("/favicon.ico", "favicon.ico")

	e.GET("/", h.GetIndex)
	e.GET("/home", h.GetHome)
	e.POST("/signup", h.SignUp)
	e.POST("/login", h.Login)

	r := e.Group("/restricted")
	r.GET("/users", h.GetUser)
	r.PUT("/users", h.UpdateUser)
	r.POST("/users", h.CreateUser)
	logger, _ := zap.NewProduction()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.Time("time", v.StartTime),
				zap.String("id", v.RequestID),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("host", v.Host),
				zap.String("method", v.Method),
				zap.String("URI", v.URI),
				zap.String("user_agent", v.UserAgent),
				zap.Int("status", v.Status),
				zap.Error(v.Error),
				zap.Duration("latency", v.Latency),
				zap.Any("headers", v.Headers),
			)

			return nil
		},
	}))

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
