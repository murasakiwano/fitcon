package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/murasakiwano/fitcon/internal/auth"
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
	e.GET("/signup", h.GetSignUp)
	e.POST("/signup", h.SignUp)
	e.GET("/login", h.GetLogin)
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

	a := e.Group("/admin")
	a.POST("", h.CreateAdmin)
	a.GET("", h.GetSignUpAdmin)
	a.POST("/login", h.LoginAdmin)
	a.GET("/login", h.GetLoginAdmin)

	e.Use(authMiddleware)

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SessionName, c)
		c.Logger().Debug(sess, err)
		if err != nil {
			return err
		}

		authenticated := sess.Values["authenticated"] == true
		requestIsToLogin := c.Request().URL.Path == "/login"
		requestMethodIsGet := c.Request().Method == http.MethodGet
		if !authenticated && !requestIsToLogin && !requestMethodIsGet {
			c.Logger().Debug(sess, err)
			c.Logger().Info("redirecting to login page...")
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		c.Logger().Debug(sess, err)

		return next(c)
	}
}

func restrictedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SessionName, c)
		c.Logger().Debug(sess, err)
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
