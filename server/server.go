package server

import (
	"2022_2_GoTo_team/server/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func routing(e *echo.Echo) {
	api := api.GetApi()
	//e.GET("/", api.RootHandler)
	//e.POST("/login", api.LoginHandler)
	//e.POST("/logout", api.LogoutHandler)
	//e.GET("/user", api.UserHandler)

	e.POST("/api/v1/user/signup", api.SignupUserHandler)
	e.POST("/api/v1/session/create", api.CreateSessionHandler)
	e.GET("/api/v1/feed", api.FeedHandler)
}

func Run(serverAddress string, allowOriginsAddressesCORS []string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowOrigins:     allowOriginsAddressesCORS,
			AllowCredentials: true,
		},
	))

	routing(e)
	e.Logger.Fatal(e.Start(serverAddress))
}
