package repositoriesConnectionsUtils

import (
	"2022_2_GoTo_team/pkg/utils/logger"
	"github.com/sirupsen/logrus"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
)

func TestRepositoriesConnectionsUtilsNegative(t *testing.T) {
	/*
		_ = GetPostgreSQLConnections("dbUser", "dbName", "dbPassword", "dbHost", "8000", "10", &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
			"requestId": "asd",
			"userEmail": "asd",
		})})
	*/
	res := GetGrpcServiceConnection("127.0.0.1:-1111", &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})
	if res == nil {
		t.Error("res is nil")
	}
}
