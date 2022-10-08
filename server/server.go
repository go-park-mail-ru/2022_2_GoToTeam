package server

import (
	"2022_2_GoTo_team/server/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routing(e *echo.Echo) {
	ap := api.GetApi()
	e.GET("/", ap.RootHandler)
	e.POST("/login", ap.LoginHandler)
	e.POST("/logout", ap.LogoutHandler)

	e.POST("/api/v1/user/signup", ap.SignupUserHandler)
	e.OPTIONS("/api/v1/user/signup", ap.SignupUserHandler)

	e.POST("/api/v1/session/create", ap.CreateSessionHandler)
	e.OPTIONS("/api/v1/session/create", ap.CreateSessionHandler)

	e.GET("/api/v1/feed", ap.FeedHandler)
	e.OPTIONS("/api/v1/feed", ap.FeedHandler)
}

func Run(servAddress string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"POST", "GET", "OPTIONS"},
	}))

	routing(e)
	e.Logger.Fatal(e.Start(servAddress))
}
