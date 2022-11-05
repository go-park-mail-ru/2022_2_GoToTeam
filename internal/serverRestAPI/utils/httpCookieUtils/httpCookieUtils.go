package httpCookieUtils

import (
	"net/http"
	"time"
)

const SESSION_HEADER_NAME = "session_id"

func MakeHttpCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:  SESSION_HEADER_NAME,
		Path:  "/",
		Value: sessionId,
		// HttpOnly: true,
		Expires: time.Now().Add(23 * time.Hour), // Note! Change value in ExpireHttpCookie function if you change hours
	}
}

func ExpireHttpCookie(httpCookie *http.Cookie) {
	httpCookie.Expires = time.Now().AddDate(0, 0, -1) // Note! Change value in MakeHttpCookie if you change days
}
