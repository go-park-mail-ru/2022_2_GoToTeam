package serverRestAPI

import (
	articleComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery"
	articleComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/articleComponent/repository"
	articleComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/articleComponent/usecase"
	categoryComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/delivery"
	categoryComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/repository"
	categoryComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/usecase"
	commentaryComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery"
	commentaryComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/repository"
	commentaryComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	feedComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery"
	feedComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/repository"
	feedComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/usecase"
	fileComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/fileComponent/delivery"
	fileComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/fileComponent/repository"
	fileComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/fileComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/middleware"
	profileComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/profileComponent/delivery"
	profileComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/profileComponent/repository"
	profileComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/profileComponent/usecase"
	searchComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/searchComponent/delivery"
	searchComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/searchComponent/repository"
	searchComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/searchComponent/usecase"
	sessionComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/delivery"
	sessionComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/repository"
	sessionComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/usecase"
	tagComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/tagComponent/delivery"
	tagComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/tagComponent/repository"
	tagComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/tagComponent/usecase"
	userComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery"
	userComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/userComponent/repository"
	userComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/userComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/configReader"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"2022_2_GoTo_team/pkg/utils/repositoriesConnectionsUtils"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	globalLogger     *logger.Logger
	middlewareLogger *logger.Logger
)

func Run(configFilePath string) {
	config, err := configReader.NewConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config settings: ")
	log.Println(config)

	globalLogger, err = logger.NewLogger(config.LogLevel, config.LogFilePath)
	if err != nil {
		log.Fatal(err)
	}
	middlewareLogger = globalLogger.ConfigureLogger("middlewareComponent", domain.LAYER_MIDDLEWARE_STRING_FOR_LOGGER)

	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(
		echoMiddleware.CORSConfig{
			AllowOrigins:     config.AllowOriginsAddressesCORS,
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowCredentials: true,
		},
	))

	if config.EnableEchoCsrfToken {
		e.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
			Skipper:        echoMiddleware.DefaultSkipper,
			TokenLength:    32,
			TokenLookup:    "header:X-XSRF-Token",
			ContextKey:     "csrf",
			CookieName:     "_csrf",
			CookieMaxAge:   82800,
			CookieSameSite: http.SameSiteDefaultMode,
			CookiePath:     "/",
			CookieHTTPOnly: false,
			CookieSecure:   false,
		}))

		middlewareLogger.LogrusLogger.Info("Echo CSRF Token security enabled.")
	} else {
		middlewareLogger.LogrusLogger.Info("Echo CSRF Token security disabled.")
	}

	if config.EnableEchoSecurity {
		e.Use(echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
			Skipper:            echoMiddleware.DefaultSkipper,
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
			XFrameOptions:      "SAMEORIGIN",
		}))

		middlewareLogger.LogrusLogger.Info("Echo security enabled.")
	} else {
		middlewareLogger.LogrusLogger.Info("Echo security disabled.")
	}

	//e.Use(echoMiddleware.Recover())
	e.Use(middleware.PanicRestoreMiddleware(middlewareLogger))
	e.Use(middleware.AccessLogMiddleware(middlewareLogger, config.EnableEchoCsrfToken))
	middleware.RegisterPrometheusMetrics()
	e.Any("/metrics", echo.WrapHandler(promhttp.Handler()))

	if err := configureServer(e, config); err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while configuring server", err))
	}
	if err := e.Start(config.ServerAddress); err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while starting server", err))
	}
}

func configureServer(e *echo.Echo, config *configReader.Config) error {
	// Loggers
	sessionDeliveryLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	sessionUsecaseLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	sessionRepositoryLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	userDeliveryLogger := globalLogger.ConfigureLogger("userComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	userUsecaseLogger := globalLogger.ConfigureLogger("userComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	userRepositoryLogger := globalLogger.ConfigureLogger("userComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	feedDeliveryLogger := globalLogger.ConfigureLogger("feedComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	feedUsecaseLogger := globalLogger.ConfigureLogger("feedComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	feedRepositoryLogger := globalLogger.ConfigureLogger("feedComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	categoryDeliveryLogger := globalLogger.ConfigureLogger("categoryComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	categoryUsecaseLogger := globalLogger.ConfigureLogger("categoryComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	categoryRepositoryLogger := globalLogger.ConfigureLogger("categoryComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	articleDeliveryLogger := globalLogger.ConfigureLogger("articleComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	articleUsecaseLogger := globalLogger.ConfigureLogger("articleComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	articleRepositoryLogger := globalLogger.ConfigureLogger("articleComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	profileDeliveryLogger := globalLogger.ConfigureLogger("profileComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	profileUsecaseLogger := globalLogger.ConfigureLogger("profileComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	profileRepositoryLogger := globalLogger.ConfigureLogger("profileComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	fileDeliveryLogger := globalLogger.ConfigureLogger("fileComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	fileUsecaseLogger := globalLogger.ConfigureLogger("fileComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	fileRepositoryLogger := globalLogger.ConfigureLogger("fileComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	tagDeliveryLogger := globalLogger.ConfigureLogger("tagComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	tagUsecaseLogger := globalLogger.ConfigureLogger("tagComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	tagRepositoryLogger := globalLogger.ConfigureLogger("tagComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	searchDeliveryLogger := globalLogger.ConfigureLogger("searchComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	searchUsecaseLogger := globalLogger.ConfigureLogger("searchComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	searchRepositoryLogger := globalLogger.ConfigureLogger("searchComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	commentaryDeliveryLogger := globalLogger.ConfigureLogger("commentaryComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	commentaryUsecaseLogger := globalLogger.ConfigureLogger("commentaryComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	commentaryRepositoryLogger := globalLogger.ConfigureLogger("commentaryComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	// PostgreSQL connections
	postgreSQLConnections := repositoriesConnectionsUtils.GetPostgreSQLConnections(config.DatabaseUser, config.DatabaseName, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseMaxOpenConnections, middlewareLogger)

	// AuthSessionService connection
	authSessionServiceConnection := repositoriesConnectionsUtils.GetGrpcServiceConnection(config.AuthSessionServiceAddress, middlewareLogger)
	// UserProfileService connection
	userProfileServiceConnection := repositoriesConnectionsUtils.GetGrpcServiceConnection(config.UserProfileServiceAddress, middlewareLogger)

	// Repositories
	sessionRepository := sessionComponentRepository.NewAuthSessionServiceRepository(authSessionServiceConnection, sessionRepositoryLogger)
	userRepository := userComponentRepository.NewUserPostgreSQLRepository(postgreSQLConnections, userRepositoryLogger)
	feedRepository := feedComponentRepository.NewFeedPostgreSQLRepository(postgreSQLConnections, feedRepositoryLogger)
	categoryRepository := categoryComponentRepository.NewCategoryPostgreSQLRepository(postgreSQLConnections, categoryRepositoryLogger)
	articleRepository := articleComponentRepository.NewArticlePostgreSQLRepository(postgreSQLConnections, articleRepositoryLogger)
	profileRepository := profileComponentRepository.NewUserProfileServiceRepository(userProfileServiceConnection, profileRepositoryLogger)
	fileRepository := fileComponentRepository.NewFileRepository(postgreSQLConnections, config.StaticDirAbsolutePath, config.ProfilePhotosDirPath, fileRepositoryLogger)
	tagRepository := tagComponentRepository.NewTagPostgreSQLRepository(postgreSQLConnections, tagRepositoryLogger)
	searchRepository := searchComponentRepository.NewSearchPostgreSQLRepository(postgreSQLConnections, searchRepositoryLogger)
	commentaryRepository := commentaryComponentRepository.NewCommentaryPostgreSQLRepository(postgreSQLConnections, commentaryRepositoryLogger)

	// Usecases and Deliveries
	sessionUsecase := sessionComponentUsecase.NewSessionUsecase(sessionRepository, sessionUsecaseLogger)
	sessionController := sessionComponentDelivery.NewSessionController(sessionUsecase, sessionDeliveryLogger)

	userUsecase := userComponentUsecase.NewUserUsecase(userRepository, userUsecaseLogger)
	userController := userComponentDelivery.NewUserController(userUsecase, sessionUsecase, userDeliveryLogger)

	feedUsecase := feedComponentUsecase.NewFeedUsecase(feedRepository, feedUsecaseLogger)
	feedController := feedComponentDelivery.NewFeedController(feedUsecase, feedDeliveryLogger)

	categoryUsecase := categoryComponentUsecase.NewCategoryUsecase(categoryRepository, categoryUsecaseLogger)
	categoryController := categoryComponentDelivery.NewCategoryController(categoryUsecase, categoryDeliveryLogger)

	articleUsecase := articleComponentUsecase.NewArticleUsecase(articleRepository, articleUsecaseLogger)
	articleController := articleComponentDelivery.NewArticleController(articleUsecase, articleDeliveryLogger)

	profileUsecase := profileComponentUsecase.NewProfileUsecase(profileRepository, profileUsecaseLogger)
	profileController := profileComponentDelivery.NewProfileController(profileUsecase, profileDeliveryLogger)

	fileUsecase := fileComponentUsecase.NewFileUsecase(fileRepository, fileUsecaseLogger)
	fileController := fileComponentDelivery.NewFileController(fileUsecase, fileDeliveryLogger)

	tagUsecae := tagComponentUsecase.NewTagUsecase(tagRepository, tagUsecaseLogger)
	tagController := tagComponentDelivery.NewTagController(tagUsecae, tagDeliveryLogger)

	searchUsecase := searchComponentUsecase.NewSearchUsecase(searchRepository, searchUsecaseLogger)
	searchController := searchComponentDelivery.NewSearchController(searchUsecase, searchDeliveryLogger)

	commentaryUsecase := commentaryComponentUsecase.NewCommentaryUsecase(commentaryRepository, commentaryUsecaseLogger)
	commentaryController := commentaryComponentDelivery.NewCommentaryController(commentaryUsecase, commentaryDeliveryLogger)

	e.Use(middleware.AuthMiddleware(sessionUsecase, middlewareLogger)) // Auth Middleware

	e.POST("/api/v1/session/create", sessionController.CreateSessionHandler)
	e.POST("/api/v1/session/remove", sessionController.RemoveSessionHandler)
	e.GET("/api/v1/session/info", sessionController.SessionInfoHandler)

	e.POST("/api/v1/user/signup", userController.SignupUserHandler)
	e.GET("/api/v1/user/info", userController.UserInfoHandler)
	e.POST("/api/v1/user/subscribe", userController.SubscribeHandler)
	e.POST("/api/v1/user/unsubscribe", userController.UnsubscribeHandler)
	e.GET("/api/v1/user/avatar", userController.GetUserAvatar)

	e.GET("/api/v1/category/info", categoryController.CategoryInfoHandler)
	e.GET("/api/v1/category/list", categoryController.CategoryListHandler)
	e.POST("/api/v1/category/subscribe", categoryController.SubscribeHandler)
	e.POST("/api/v1/category/unsubscribe", categoryController.UnsubscribeHandler)

	e.GET("/api/v1/feed", feedController.FeedHandler)
	e.GET("/api/v1/feed/user", feedController.FeedUserHandler)
	e.GET("/api/v1/feed/category", feedController.FeedCategoryHandler)

	e.GET("/api/v1/article", articleController.ArticleHandler)
	e.POST("/api/v1/article/create", articleController.CreateArticleHandler)
	e.POST("/api/v1/article/remove", articleController.RemoveArticleHandler)
	e.POST("/api/v1/article/update", articleController.UpdateArticleHandler)

	e.GET("/api/v1/profile", profileController.GetProfileHandler)
	e.POST("/api/v1/profile/update", profileController.UpdateProfileHandler)

	e.POST("/api/v1/file/upload/profile/photo", fileController.UploadProfilePhotoHandler)

	e.GET("/api/v1/tag/list", tagController.TagsListHandler)

	e.GET("/api/v1/search", searchController.SearchHandler)
	e.GET("/api/v1/search/tag", searchController.SearchTagHandler)

	e.POST("/api/v1/commentary/create", commentaryController.CreateCommentaryHandler)
	e.GET("/api/v1/commentary/feed", commentaryController.GetAllCommentariesForArticle)

	return nil
}
