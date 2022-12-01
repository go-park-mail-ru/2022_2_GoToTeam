package repositoriesConnectionsUtils

import (
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"database/sql"
	"google.golang.org/grpc"
	"strconv"
)

func GetPostgreSQLConnections(databaseUser string, databaseName string, databasePassword string, databaseHost string, databasePort string, databaseMaxOpenConnections string, logger *logger.Logger) *sql.DB {
	dsn := "user=" + databaseUser + " dbname=" + databaseName + " password=" + databasePassword + " host=" + databaseHost + " port=" + databasePort + " sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.LogrusLogger.Fatal(errorsUtils.WrapError("error while opening connection to database", err))
	}
	// Test connection
	if err = db.Ping(); err != nil {
		logger.LogrusLogger.Fatal(errorsUtils.WrapError("error while testing connection to database", err))
	}

	databaseMaxOpenConnectionsINT, err := strconv.Atoi(databaseMaxOpenConnections)
	if err != nil {
		logger.LogrusLogger.Fatal(errorsUtils.WrapError("error while parsing databaseMaxOpenConnections to int value", err))
	}
	db.SetMaxOpenConns(databaseMaxOpenConnectionsINT)

	return db
}

func GetGrpcServiceConnection(serverAddress string, logger *logger.Logger) *grpc.ClientConn {
	logger.LogrusLogger.Debug("Enter to the getGrpcServiceConnection function.")

	grcpConn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	// defer grcpConn.Close()

	if err != nil {
		logger.LogrusLogger.Fatal(errorsUtils.WrapError("error while opening connection to grpc service", err))
	}

	return grcpConn
}
