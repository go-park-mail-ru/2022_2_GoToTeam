package models

import "net/http"

type Session struct {
	Cookie *http.Cookie
}
