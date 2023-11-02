package handler

import (
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/murasakiwano/fitcon/components"
	"go.uber.org/zap"
)

var (
	key   = os.Getenv("JWT_SECRET")
	store = sessions.NewCookieStore([]byte(key))
)

func (h *Handler) Serve() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
	e.Use(authMiddleware)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			sugar.Infow("request values",
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

	e.Static("/assets", "assets")

	e.File("/favicon.ico", "favicon.ico")

	e.GET("/", h.GetIndex)
	e.GET("/signup", h.GetSignUp)
	e.POST("/signup", h.SignUp)
	e.GET("/login", h.GetLogin)
	e.POST("/login", h.Login)

	r := e.Group("/restricted")
	r.Use(restrictedMiddleware)
	r.GET("/users", h.GetUser)

	a := e.Group("/admin")
	a.POST("", h.CreateAdmin)
	a.GET("", h.GetSignUpAdmin)
	a.POST("/login", h.LoginAdmin)
	a.GET("/login", h.GetLoginAdmin)
	a.GET("/users", h.GetCreate)
	a.GET("/update_user", h.GetUpdateUser)
	a.PUT("/users", h.UpdateUser)
	a.POST("/users", h.CreateUser)

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	return func(c echo.Context) error {
		sugar.Debugw("Inside authMiddleware")
		sess, err := session.Get(SessionName, c)
		if err != nil {
			sugar.Error(err)
			return err
		}

		authenticated := sess.Values["authenticated"] == true
		path := c.Request().URL.Path
		bypassRequest := path == "/login" ||
			strings.Contains(path, "/admin") ||
			path == "/signup" ||
			strings.Contains(path, "assets")
		// requestMethodIsGet := c.Request().Method == http.MethodGet
		if authenticated || bypassRequest {
			return next(c)
		}

		sugar.Info("redirecting to login page...")
		c.Response().Header().Set("HX-Replace-Url", "/login")
		if err := components.Index(components.Login()).Render(c.Request().Context(), c.Response().Writer); err != nil {
			sugar.Error(err)
			return err
		}

		return nil
	}
}

func restrictedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SessionName, c)
		c.Logger().Debugf("got session: %+v", sess)
		if err != nil {
			return err
		}

		id := c.QueryParam("matricula")
		admin := sess.Values["admin"] == true
		validID := sess.Values["user_id"] == id
		authorized := admin || validID
		c.Logger().Debugj(log.JSON{
			"admin":   admin,
			"validID": validID,
			"user_id": sess.Values["user_id"],
			"id":      id,
		})
		if !authorized {
			return echo.ErrForbidden
		}

		return next(c)
	}
}
