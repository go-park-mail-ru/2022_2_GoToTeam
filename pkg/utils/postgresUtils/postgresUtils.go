package postgresUtils

import (
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"database/sql"
	"strconv"
)

func GetPostgreSQLConnections(databaseUser string, databaseName string, databasePassword string, databaseHost string, databasePort string, databaseMaxOpenConnections string, middlewareLogger *logger.Logger) *sql.DB {
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
