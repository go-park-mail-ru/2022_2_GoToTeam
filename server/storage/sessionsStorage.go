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

func (o *SessionsStorage) CreateSessionForUser(email string) *http.Cookie {
	SID := randStringRunes(32)

	o.sessions[SID] = email

	return &http.Cookie{
		Name:  "session_id",
		Path:  "/",
		Value: SID,
		// HttpOnly: true,
		Expires: time.Now().Add(23 * time.Hour), // Note! Change value in cookie.Expires function if you change hours
	}
}

func (o *SessionsStorage) RemoveSession(cookie *http.Cookie) {
	delete(o.sessions, cookie.Value)
	o.ExpireCookie(cookie)
}

func (o *SessionsStorage) SessionExists(cookie string) bool {
	_, exists := o.sessions[cookie]
	return exists
}

func (o *SessionsStorage) ExpireCookie(cookie *http.Cookie) {
	cookie.Expires = time.Now().AddDate(0, 0, -1) // Note! Change value in create cookie expires if you change days
}
