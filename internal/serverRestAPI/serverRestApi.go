package serverRestAPI

import (
	"2022_2_GoTo_team/internal/serverRestAPI/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func routing(e *echo.Echo) {
	api := api.GetApi()
	e.POST("/api/v1/session/create", api.CreateSessionHandler)
	e.POST("/api/v1/session/remove", api.RemoveSessionHandler)
	e.GET("/api/v1/session/info", api.SessionInfoHandler)

	e.POST("/api/v1/new/article/create", api.CreateArticleHandler)
	e.POST("/api/v1/new/article/update", api.UpdateArticleHandler)

	e.POST("/api/v1/user/signup", api.SignupUserHandler)
	e.GET("/api/v1/user/info", api.UserInfoHandler)
	e.GET("/api/v1/user/feed", api.UserFeedHandler)

	e.GET("/api/v1/category/info", api.CategoryInfoHandler)
	e.GET("/api/v1/category/feed", api.CategoryFeedHandler)

	e.GET("/api/v1/feed", api.FeedHandler)

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

	routing(e)
	e.Logger.Fatal(e.Start(conf.BindServerAddress))
}
