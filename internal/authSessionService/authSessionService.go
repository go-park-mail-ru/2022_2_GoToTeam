package authSessionService

import (
	"2022_2_GoTo_team/internal/authSessionService/domain"
	"2022_2_GoTo_team/internal/authSessionService/middleware"
	sessionComponentDelivery "2022_2_GoTo_team/internal/authSessionService/sessionComponent/delivery"
	sessionComponentRepository "2022_2_GoTo_team/internal/authSessionService/sessionComponent/repository"
	sessionComponentUsecase "2022_2_GoTo_team/internal/authSessionService/sessionComponent/usecase"
	userComponentRepository "2022_2_GoTo_team/internal/authSessionService/userComponent/repository"
	"2022_2_GoTo_team/internal/authSessionService/utils/configReader"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
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
	sessionDeliveryLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_DELIVERY_STRING_FOR_LOGGER)
	sessionUsecaseLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_USECASE_STRING_FOR_LOGGER)
	sessionRepositoryLogger := globalLogger.ConfigureLogger("sessionComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)
	userRepositoryLogger := globalLogger.ConfigureLogger("userComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER)

	// PostgreSQL connections
	postgreSQLConnections := repositoriesConnectionsUtils.GetPostgreSQLConnections(config.DatabaseUser, config.DatabaseName, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseMaxOpenConnections, middlewareLogger)

	// Repositories
	sessionRepository := sessionComponentRepository.NewSessionCustomRepository(sessionRepositoryLogger)
	userRepository := userComponentRepository.NewUserPostgreSQLRepository(postgreSQLConnections, userRepositoryLogger)

	// Usecases and Deliveries
	sessionUsecase := sessionComponentUsecase.NewSessionUsecase(sessionRepository, userRepository, sessionUsecaseLogger)
	sessionDelivery := sessionComponentDelivery.NewSessionDelivery(sessionUsecase, sessionDeliveryLogger)

	authSessionServiceGrpcProtos.RegisterAuthSessionServiceServer(server, sessionDelivery)

	if err := server.Serve(listener); err != nil {
		middlewareLogger.LogrusLogger.Fatal(errorsUtils.WrapError("error while serve", err))
	}

	middlewareLogger.LogrusLogger.Info("grpc server started on ", config.ServerAddress)
}
