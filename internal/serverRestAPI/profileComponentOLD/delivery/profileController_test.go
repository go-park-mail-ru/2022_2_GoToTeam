package delivery

/*
import "database/sql"


import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	profileComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/profileComponentOLD/repository"
	profileComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/profileComponentOLD/usecase"
	sessionComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/sessionComponentOLD/repository"
	sessionComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/sessionComponentOLD/usecase"
	userComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/userComponent/repository"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	logger2 "2022_2_GoTo_team/pkg/utils/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func getPostgreSQLConnections(databaseUser string, databaseName string, databasePassword string, databaseHost string, databasePort string, databaseMaxOpenConnections string, middlewareLogger *logger2.Logger) *sql.DB {
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

// Создает папку с логами не там где нужно
func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()

	// подготавливаем правильный ответ
	profileResp := new(models.Profile)
	profileResp.Email = "admin@admin.admin"
	profileResp.Login = "admin"
	profileResp.Username = "Mikle"
	profileResp.AvatarImgPath = "/avatars/admin.jpeg"
	res, _ := json.Marshal(profileResp)
	fmt.Println("RES:::")
	fmt.Println(res)
	//loginResponse := new(models.LoginResponse)
	//loginResponse.Status = 200
	//loginResponse.Data = *profileResp
	//loginResponse.Msg = "OK"
	//loginRes, _ := json.Marshal(loginResponse)

	//подготавливаем тест
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(req, recorder)
	c.SetPath("/feed")

	profileDeliveryLogger, _ := logger2.GetConfiguredLoggerWrapper("profileComponentOLD", domain.LAYER_DELIVERY_STRING_FOR_LOGGER, "debug", "logs/serverRestApi/logs.log")
	profileUsecaseLogger, _ := logger2.GetConfiguredLoggerWrapper("profileComponentOLD", domain.LAYER_USECASE_STRING_FOR_LOGGER, "debug", "logs/serverRestApi/logs.log")
	sessionUsecaseLogger, _ := logger2.GetConfiguredLoggerWrapper("sessionComponentOLD", domain.LAYER_USECASE_STRING_FOR_LOGGER, "debug", "logs/serverRestApi/logs.log")
	sessionRepositoryLogger, _ := logger2.GetConfiguredLoggerWrapper("sessionComponentOLD", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, "debug", "logs/serverRestApi/logs.log")
	profileRepositoryLogger, _ := logger2.GetConfiguredLoggerWrapper("profileComponentOLD", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, "debug", "logs/serverRestApi/logs.log")
	userRepositoryLogger, _ := logger2.GetConfiguredLoggerWrapper("userComponent", domain.LAYER_REPOSITORY_STRING_FOR_LOGGER, "debug", "logs/serverRestApi/logs.log")

	middlewareLogger, _ := logger2.GetConfiguredLoggerWrapper("middlewareComponent", "middlewareLayer", "debug", "logs/serverRestApi/logs.log")
	postgreSQLConnections := getPostgreSQLConnections("postgres", "ve_ru", "admin", "127.0.0.1", "5432", "10", middlewareLogger)
	userRepository := userComponentRepository.NewUserPostgreSQLRepository(postgreSQLConnections, userRepositoryLogger)
	profileRepository := profileComponentRepository.NewProfilePostgreSQLRepository(postgreSQLConnections, profileRepositoryLogger)
	sessionRepository := sessionComponentRepository.NewSessionCustomRepository(sessionRepositoryLogger)
	sessionUsecase := sessionComponentUsecase.NewSessionUsecase(sessionRepository, userRepository, sessionUsecaseLogger)
	profileUsecase := profileComponentUsecase.NewProfileUsecase(profileRepository, sessionRepository, profileUsecaseLogger)
	controllerToTest := NewProfileController(profileUsecase, sessionUsecase, profileDeliveryLogger)

	// вызываем тест + Assertions
	if assert.NoError(t, controllerToTest.GetProfileHandler(c)) {
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Equal(t, "", recorder.Body.String())
	}
}

*/
