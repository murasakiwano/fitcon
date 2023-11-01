package handler

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/murasakiwano/fitcon/components"
	"github.com/murasakiwano/fitcon/internal/auth"
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
	e.Static("/img", "img")

	e.File("/favicon.ico", "favicon.ico")

	e.GET("/", h.GetIndex)
	e.GET("/signup", h.GetSignUp)
	e.POST("/signup", h.SignUp)
	e.GET("/login", h.GetLogin)
	e.POST("/login", h.Login)

	r := e.Group("/restricted")
	r.GET("/users", h.GetUser)
	r.PUT("/users", h.UpdateUser)
	r.POST("/users", h.CreateUser)
	r.Use(restrictedMiddleware)

	a := e.Group("/admin")
	a.POST("", h.CreateAdmin)
	a.GET("", h.GetSignUpAdmin)
	a.POST("/login", h.LoginAdmin)
	a.GET("/login", h.GetLoginAdmin)

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return func(c echo.Context) error {
		sess, err := session.Get(SessionName, c)
		if err != nil {
			sugar.Error(err)
			return err
		}

		authenticated := sess.Values["authenticated"] == true
		path := c.Request().URL.Path
		requestIsToLogin := path == "/login" || path == "/admin"
		// requestMethodIsGet := c.Request().Method == http.MethodGet
		if authenticated || requestIsToLogin {
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
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return func(c echo.Context) error {
		sess, err := session.Get(SessionName, c)
		sugar.Debugw("values", sess, err)
		if err != nil {
			return err
		}

		id := c.QueryParam("id")
		tokenString := sess.Values["token"]
		jwtInfo, err := auth.ValidateJWT(
			tokenString.(string),
			os.Getenv("JWT_SECRET"),
		)
		if err != nil {
			return err
		}

		authorized := jwtInfo.Issuer == "fitcon" &&
			(jwtInfo.Subject == id || jwtInfo.Admin)
		if !authorized {
			return echo.ErrForbidden
		}

		return next(c)
	}
}
