package serverRestAPI

import (
	articleComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery"
	articleComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/articleComponent/repository"
	articleComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/articleComponent/usecase"
	categoryComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/delivery"
	categoryComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/repository"
	categoryComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	feedComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery"
	feedComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/repository"
	feedComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/middleware"
	sessionComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/delivery"
	sessionComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/repository"
	sessionComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/usecase"
	userComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery"
	userComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/userComponent/repository"
	userComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/userComponent/usecase"

	profileComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/profileComponent/delivery"
	profileComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/profileComponent/repository"
	profileComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/profileComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"database/sql"
	"net/http"
	"strconv"

	"2022_2_GoTo_team/internal/serverRestAPI/utils/configReader"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"log"
)

func Run(configFilePath string) {
	config, err := configReader.NewConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config settings: ")
	log.Println(config)

	middlewareLogger, err := logger.NewLogger("middlewareComponent", domain.LAYER_MIDDLEWARE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(
		echoMiddleware.CORSConfig{
			AllowOrigins:     config.AllowOriginsAddressesCORS,
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowCredentials: true,
		},
	))

	//e.Use(echoMiddleware.Recover())
	e.Use(middleware.PanicRestoreMiddleware(middlewareLogger))
	e.Use(middleware.AccessLogMiddleware(middlewareLogger))

	if err := configureServer(e, config, middlewareLogger); err != nil {
		middlewareLogger.LogrusLogger.Error(errorsUtils.WrapError("error while configuring server", err))
		e.Logger.Fatal(errorsUtils.WrapError("error while configuring server", err))
	}

	if err := e.Start(config.ServerAddress); err != nil {
		middlewareLogger.LogrusLogger.Error(errorsUtils.WrapError("error while starting server", err))
		e.Logger.Fatal(errorsUtils.WrapError("error while starting server", err))
	}
}

func getPostgreSQLConnections(databaseUser string, databaseName string, databasePassword string, databaseHost string, databasePort string, databaseMaxOpenConnections string, middlewareLogger *logger.Logger) *sql.DB {
	dsn := "user=" + databaseUser + " dbname=" + databaseName + " password=" + databasePassword + " host=" + databaseHost + " port=" + databasePort + " sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while opening connection to database", err))
	}
	// Test connection
	if err = db.Ping(); err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while testing connection to database", err))
	}

	databaseMaxOpenConnectionsINT, err := strconv.Atoi(databaseMaxOpenConnections)
	if err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while parsing databaseMaxOpenConnections to int value", err))
	}
	db.SetMaxOpenConns(databaseMaxOpenConnectionsINT)

	return db
}

func configureServer(e *echo.Echo, config *configReader.Config, middlewareLogger *logger.Logger) error {

	// Loggers
	sessionDeliveryLogger, err := logger.NewLogger("sessionComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionUsecaseLogger, err := logger.NewLogger("sessionComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionRepositoryLogger, err := logger.NewLogger("sessionComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userDeliveryLogger, err := logger.NewLogger("userComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userUsecaseLogger, err := logger.NewLogger("userComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userRepositoryLogger, err := logger.NewLogger("userComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	feedDeliveryLogger, err := logger.NewLogger("feedComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	feedUsecaseLogger, err := logger.NewLogger("feedComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	feedRepositoryLogger, err := logger.NewLogger("feedComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	categoryDeliveryLogger, err := logger.NewLogger("categoryComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	categoryUsecaseLogger, err := logger.NewLogger("categoryComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	categoryRepositoryLogger, err := logger.NewLogger("categoryComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	articleDeliveryLogger, err := logger.NewLogger("articleComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	articleUsecaseLogger, err := logger.NewLogger("articleComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	articleRepositoryLogger, err := logger.NewLogger("articleComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	profileDeliveryLogger, err := logger.NewLogger("profileComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	profileUsecaseLogger, err := logger.NewLogger("profileComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	profileRepositoryLogger, err := logger.NewLogger("profileComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}

	// PostgreSQL connections
	postgreSQLConnections := getPostgreSQLConnections(config.DatabaseUser, config.DatabaseName, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseMaxOpenConnections, middlewareLogger)

	// Repositories
	sessionRepository := sessionComponentRepository.NewSessionCustomRepository(sessionRepositoryLogger)
	userRepository := userComponentRepository.NewUserPostgreSQLRepository(postgreSQLConnections, userRepositoryLogger)
	feedRepository := feedComponentRepository.NewFeedPostgreSQLRepository(postgreSQLConnections, feedRepositoryLogger)
	categoryRepository := categoryComponentRepository.NewCategoryPostgreSQLRepository(postgreSQLConnections, categoryRepositoryLogger)
	articleRepository := articleComponentRepository.NewArticlePostgreSQLRepository(postgreSQLConnections, articleRepositoryLogger)
	profileRepository := profileComponentRepository.NewProfilePostgreSQLRepository(postgreSQLConnections, profileRepositoryLogger)
	// Usecases and Deliveries
	sessionUsecase := sessionComponentUsecase.NewSessionUsecase(sessionRepository, userRepository, sessionUsecaseLogger)
	sessionController := sessionComponentDelivery.NewSessionController(sessionUsecase, sessionDeliveryLogger)

	userUsecase := userComponentUsecase.NewUserUsecase(userRepository, userUsecaseLogger)
	userController := userComponentDelivery.NewUserController(userUsecase, sessionUsecase, userDeliveryLogger)

	feedUsecase := feedComponentUsecase.NewFeedUsecase(feedRepository, feedUsecaseLogger)
	feedController := feedComponentDelivery.NewFeedController(feedUsecase, feedDeliveryLogger)

	categoryUsecase := categoryComponentUsecase.NewCategoryUsecase(categoryRepository, categoryUsecaseLogger)
	categoryController := categoryComponentDelivery.NewCategoryController(categoryUsecase, categoryDeliveryLogger)

	articleUsecase := articleComponentUsecase.NewArticleUsecase(articleRepository, sessionRepository, articleUsecaseLogger)
	articleController := articleComponentDelivery.NewArticleController(articleUsecase, articleDeliveryLogger)

	profileUsecase := profileComponentUsecase.NewProfileUsecase(profileRepository, sessionRepository, profileUsecaseLogger)
	profileController := profileComponentDelivery.NewProfileController(profileUsecase, sessionUsecase, profileDeliveryLogger)

	e.Use(middleware.AuthMiddleware(sessionUsecase, middlewareLogger)) // Auth Middleware

	e.POST("/api/v1/session/create", sessionController.CreateSessionHandler)
	e.POST("/api/v1/session/remove", sessionController.RemoveSessionHandler)
	e.GET("/api/v1/session/info", sessionController.SessionInfoHandler)

	e.POST("/api/v1/user/signup", userController.SignupUserHandler)
	e.GET("/api/v1/user/info", userController.UserInfoHandler)

	e.GET("/api/v1/category/info", categoryController.CategoryInfoHandler)

	e.GET("/api/v1/feed", feedController.FeedHandler)
	e.GET("/api/v1/feed/user", feedController.FeedUserHandler)
	e.GET("/api/v1/feed/category", feedController.FeedCategoryHandler)

	e.GET("/api/v1/article", articleController.ArticleHandler)
	e.POST("/api/v1/article/create", articleController.ArticleCreateHandler)

	e.GET("/api/v1/profile", profileController.GetProfileHandler)
	e.POST("/api/v1/profile/update", profileController.UpdateProfileHandler)

	return nil
}
