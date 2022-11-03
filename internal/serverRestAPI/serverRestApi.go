package serverRestAPI

import (
	"2022_2_GoTo_team/internal/serverRestAPI/api"
	"2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/delivery"
	"2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/repository"
	"2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/usecase"
	repository2 "2022_2_GoTo_team/internal/serverRestAPI/userComponent/repository"
	"2022_2_GoTo_team/internal/utils/configReader"
	"2022_2_GoTo_team/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

const (
	LAYER_DELIVERY   = "delivery"
	LAYER_USECASE    = "usecase"
	LAYER_REPOSITORY = "repository"
)

func Run(configFilePath string) {
	config, err := configReader.NewConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("config settings: ")
	log.Println(config)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     config.AllowOriginsAddressesCORS,
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowCredentials: true,
		},
	))

	if err := routing(e, config); err != nil {
		e.Logger.Fatal("cant configure logger: " + err.Error())
	}

	e.Logger.Fatal(e.Start(config.ServerAddress))
}

func routing(e *echo.Echo, config *configReader.Config) error {
	Api := api.GetApi()

	sessionComponentDeliveryLogger, err := logger.NewLogger("sessionComponent", LAYER_DELIVERY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionComponentUsecaseLogger, err := logger.NewLogger("sessionComponent", LAYER_USECASE, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionComponentRepositoryLogger, err := logger.NewLogger("sessionComponent", LAYER_REPOSITORY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userComponentRepositoryLogger, err := logger.NewLogger("userComponent", LAYER_REPOSITORY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionHandler := delivery.NewSessionHandler(
		usecase.NewSessionUsecase(
			repository.NewSessionRepository(sessionComponentRepositoryLogger),
			repository2.NewUserRepository(userComponentRepositoryLogger),
			sessionComponentUsecaseLogger,
		),
		sessionComponentDeliveryLogger,
	)

	Api.LogInfo("starting server")

	e.POST("/api/v1/session/create", sessionHandler.CreateSessionHandler)
	e.POST("/api/v1/session/remove", sessionHandler.RemoveSessionHandler)
	e.GET("/api/v1/session/info", sessionHandler.SessionInfoHandler)

	e.POST("/api/v1/article/create", Api.CreateArticleHandler)
	e.POST("/api/v1/article/update", Api.UpdateArticleHandler)

	e.POST("/api/v1/user/signup", Api.SignupUserHandler)
	e.GET("/api/v1/user/info", Api.UserInfoHandler)
	e.GET("/api/v1/user/feed", Api.UserFeedHandler)

	e.GET("/api/v1/category/info", Api.CategoryInfoHandler)
	e.GET("/api/v1/category/feed", Api.CategoryFeedHandler)

	e.GET("/api/v1/feed", Api.FeedHandler)

	return nil
}
