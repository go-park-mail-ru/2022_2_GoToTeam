package storage

import (
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

type SessionsStorage struct {
	sessions map[string]string
	mu       sync.RWMutex
}

func GetSessionsStorage() *SessionsStorage {
	return &SessionsStorage{
		sessions: make(map[string]string),
		mu:       sync.RWMutex{},
	}
}

func (o *SessionsStorage) PrintSessions() {
	log.Println("Sessions in storage:")
	for k, v := range o.sessions {
		log.Printf("cook: %#v for user email: %#v", k, v)
	}
}

func (o *SessionsStorage) CreateCookieForUser(email string) *http.Cookie {
	SID := randStringRunes(32)

	o.sessions[SID] = email

	return &http.Cookie{
		Name:  "session_id",
		Path:  "/",
		Value: SID,
		// HttpOnly: true,
		Expires: time.Now().Add(24 * time.Hour),
	}
}

func (o *SessionsStorage) SessionExists(cookie string) bool {
	_, exists := o.sessions[cookie]
	return exists
}
