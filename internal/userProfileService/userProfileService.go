package userProfileService

import (
	"2022_2_GoTo_team/internal/userProfileService/domain"
	"2022_2_GoTo_team/internal/userProfileService/middleware"
	profileComponentDelivery "2022_2_GoTo_team/internal/userProfileService/profileComponent/delivery"
	profileComponentRepository "2022_2_GoTo_team/internal/userProfileService/profileComponent/repository"
	profileComponentUsecase "2022_2_GoTo_team/internal/userProfileService/profileComponent/usecase"
	sessionComponentRepository "2022_2_GoTo_team/internal/userProfileService/sessionComponent/repository"
	"2022_2_GoTo_team/internal/userProfileService/utils/configReader"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/userProfileServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"2022_2_GoTo_team/pkg/utils/repositoriesConnectionsUtils"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc"
	"log"
	"net"
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

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while starting server", err))
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryServerInterceptor(middlewareLogger)))

	// Loggers
	profileDeliveryLogger := globalLogger.ConfigureLogger("profileComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	profileUsecaseLogger := globalLogger.ConfigureLogger("profileComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	profileRepositoryLogger := globalLogger.ConfigureLogger("profileComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)
	sessionRepositoryLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	// PostgreSQL connections
	postgreSQLConnections := repositoriesConnectionsUtils.GetPostgreSQLConnections(config.DatabaseUser, config.DatabaseName, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseMaxOpenConnections, middlewareLogger)

	// authSessionService connection
	authSessionServiceConnection := repositoriesConnectionsUtils.GetGrpcServiceConnection(config.AuthSessionServiceAddress, middlewareLogger)

	// Repositories
	profileRepository := profileComponentRepository.NewProfilePostgreSQLRepository(postgreSQLConnections, profileRepositoryLogger)
	sessionRepository := sessionComponentRepository.NewAuthSessionServiceRepository(authSessionServiceConnection, sessionRepositoryLogger)

	// Usecases and Deliveries
	profileUsecase := profileComponentUsecase.NewProfileUsecase(profileRepository, sessionRepository, profileUsecaseLogger)
	profileDelivery := profileComponentDelivery.NewProfileDelivery(profileUsecase, profileDeliveryLogger)

	userProfileServiceGrpcProtos.RegisterUserProfileServiceServer(server, profileDelivery)

	if err := server.Serve(listener); err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while serve", err))
	}

	middlewareLogger.LogrusLogger.Info("grpc server started on ", config.ServerAddress)
}
