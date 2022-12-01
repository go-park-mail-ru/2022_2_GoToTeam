package repository

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var loggerMock = &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
	"requestId": "qwerty",
	"userEmail": "asd@asd.asd",
})}

func TestCreateSessionForUser(t *testing.T) {
	scr := NewSessionCustomRepository(loggerMock)
	res, err := scr.CreateSessionForUser(context.Background(), "asd@asd.asd")
	assert.NotEqual(t, nil, res)
	assert.Equal(t, nil, err)
}

func TestGetEmailBySessionFound(t *testing.T) {
	scr := NewSessionCustomRepository(loggerMock)
	email := "asd@asd.asd"
	res, _ := scr.CreateSessionForUser(context.Background(), email)

	resultEmail, _ := scr.GetEmailBySession(context.Background(), res)

	assert.Equal(t, email, resultEmail)
}

func TestGetEmailBySessionNotFound(t *testing.T) {
	scr := NewSessionCustomRepository(loggerMock)
	resultEmail, err := scr.GetEmailBySession(context.Background(), &models.Session{
		SessionId: "sess1",
	})

	assert.Equal(t, "", resultEmail)
	assert.Equal(t, repositoryToUsecaseErrors.SessionRepositoryEmailDoesntExistError, err)
}

func TestUpdateEmailBySession(t *testing.T) {
	scr := NewSessionCustomRepository(loggerMock)
	err := scr.UpdateEmailBySession(context.Background(), &models.Session{
		SessionId: "sess1",
	}, "mewEmail@asd.asd")

	assert.Equal(t, nil, err)
}

func TestRemoveSession(t *testing.T) {
	scr := NewSessionCustomRepository(loggerMock)
	err := scr.RemoveSession(context.Background(), &models.Session{
		SessionId: "sess1",
	})

	assert.Equal(t, nil, err)
}

func TestSessionExists(t *testing.T) {
	scr := NewSessionCustomRepository(loggerMock)
	email := "asd@asd.asd"
	res, _ := scr.CreateSessionForUser(context.Background(), email)

	exists, err := scr.SessionExists(context.Background(), res)
	assert.Equal(t, true, exists)
	assert.Equal(t, nil, err)
	exists, err = scr.SessionExists(context.Background(), &models.Session{
		SessionId: "sess1",
	})
	assert.Equal(t, false, exists)
	assert.Equal(t, nil, err)

}
