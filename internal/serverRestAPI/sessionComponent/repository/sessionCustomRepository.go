package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
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

func NewSessionCustomRepository(logger *logger.Logger) sessionComponentInterfaces.SessionRepositoryInterface {
	return &sessionsStorage{
		sessions: make(map[string]string),
		mu:       sync.RWMutex{},
		logger:   logger,
	}
}

func (ss *sessionsStorage) PrintSessions() {
	log.Println("Sessions in storage:")
	for k, v := range ss.sessions {
		log.Printf("cook: %#v for user email: %#v", k, v)
	}
}

func (ss *sessionsStorage) CreateSessionForUser(email string, sessionHeaderName string) *models.Session {
	SID := randStringRunes(32)

	ss.sessions[SID] = email

	// TODO logger
	ss.PrintSessions()

	return &models.Session{
		Cookie: &http.Cookie{
			Name:  sessionHeaderName,
			Path:  "/",
			Value: SID,
			// HttpOnly: true,
			Expires: time.Now().Add(23 * time.Hour), // Note! Change value in cookie.Expires function if you change hours
		},
	}
}

func (ss *sessionsStorage) GetEmailBySession(session *models.Session) string {
	return ss.sessions[session.Cookie.Value]
}

func (ss *sessionsStorage) RemoveSession(session *models.Session) {
	delete(ss.sessions, session.Cookie.Value)
	ss.ExpireSession(session)

	// TODO logger
	ss.PrintSessions()
}

func (ss *sessionsStorage) SessionExists(session *models.Session) bool {
	_, exists := ss.sessions[session.Cookie.Value]
	return exists
}

func (ss *sessionsStorage) ExpireSession(session *models.Session) {
	session.Cookie.Expires = time.Now().AddDate(0, 0, -1) // Note! Change value in create cookie expires if you change days
}
