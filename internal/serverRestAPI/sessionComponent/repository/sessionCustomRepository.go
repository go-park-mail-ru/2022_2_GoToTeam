package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"log"
	"math/rand"
	"sync"
)

const SESSION_ID_STRING_LENGTH = 32

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomRunesString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

type sessionsStorage struct {
	sessions map[string]string
	mu       sync.RWMutex
	logger   *logger.Logger
}

func NewSessionCustomRepository(logger *logger.Logger) (sessionComponentInterfaces.SessionRepositoryInterface, error) {
	sessionsStorage := &sessionsStorage{
		sessions: make(map[string]string),
		mu:       sync.RWMutex{},
		logger:   logger,
	}

	sessionsStorage.logSessions()

	return sessionsStorage, nil
}

func (ss *sessionsStorage) logSessions() {
	// TODO logger
	log.Println("Sessions in storage:")
	for k, v := range ss.sessions {
		log.Printf("cook: %#v for user email: %#v", k, v)
	}
}

func (ss *sessionsStorage) CreateSessionForUser(email string) (*models.Session, error) {
	sessionId := generateRandomRunesString(SESSION_ID_STRING_LENGTH)

	ss.sessions[sessionId] = email

	// TODO logger
	ss.logSessions()

	return &models.Session{
		SessionId: sessionId,
	}, nil
}

func (ss *sessionsStorage) GetEmailBySession(session *models.Session) (string, error) {
	return ss.sessions[session.SessionId], nil
}

func (ss *sessionsStorage) RemoveSession(session *models.Session) error {
	delete(ss.sessions, session.SessionId)

	// TODO logger
	ss.logSessions()

	return nil
}

func (ss *sessionsStorage) SessionExists(session *models.Session) (bool, error) {
	_, exists := ss.sessions[session.SessionId]

	return exists, nil
}
