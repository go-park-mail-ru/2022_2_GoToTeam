package serverRestAPI

import (
	"2022_2_GoTo_team/internal/serverRestAPI/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func routing(e *echo.Echo, conf *Config) error {
	Api := api.GetApi()
	if err := Api.ConfigureLogger(conf.LogLevel); err != nil {
		return err
	}
	Api.LogInfo("starting server")

	e.POST("/api/v1/session/create", Api.CreateSessionHandler)
	e.POST("/api/v1/session/remove", Api.RemoveSessionHandler)
	e.GET("/api/v1/session/info", Api.SessionInfoHandler)

	e.POST("/api/v1/new/article/create", Api.CreateArticleHandler)
	e.POST("/api/v1/new/article/update", Api.UpdateArticleHandler)

	e.POST("/api/v1/user/signup", Api.SignupUserHandler)
	e.GET("/api/v1/user/info", Api.UserInfoHandler)
	e.GET("/api/v1/user/feed", Api.UserFeedHandler)

	e.GET("/api/v1/category/info", Api.CategoryInfoHandler)
	e.GET("/api/v1/category/feed", Api.CategoryFeedHandler)

	e.GET("/api/v1/feed", Api.FeedHandler)
	return nil
}

func Run(conf *Config) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     conf.BindAllowOriginsAddressesCORS,
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowCredentials: true,
		},
	))

	if err := routing(e, conf); err != nil {
		e.Logger.Fatal("Cant configure logger")
	}
	e.Logger.Fatal(e.Start(conf.BindServerAddress))
}
