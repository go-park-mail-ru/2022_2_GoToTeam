package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routing(e *echo.Echo) {
	api := api.GetApi()
	e.GET("/", api.RootHandler)
	e.POST("/login", api.LoginHandler)
	e.POST("/logout", api.LogoutHandler)
}

func Run(servAddress string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"POST", "GET"},
	}))

	routing(e)
	e.Logger.Fatal(e.Start(servAddress))
}
