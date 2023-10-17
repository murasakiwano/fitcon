package handler

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func (h *Handler) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use()

	e.Static("/assets", "assets")
	e.Static("/img", "img")
	e.Static("/css", "css")

	e.File("/favicon.ico", "favicon.ico")

	e.GET("/", h.GetIndex)
	e.GET("/users", h.GetUser)
	e.GET("/home", h.GetHome)
	e.POST("/signup", h.SignUp)
	e.POST("/login", h.Login)

	r := e.Group("/restricted")
	r.PUT("/users", h.UpdateUser)
	r.POST("/users", h.CreateUser)

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	e.Use(echojwt.WithConfig(config))

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
