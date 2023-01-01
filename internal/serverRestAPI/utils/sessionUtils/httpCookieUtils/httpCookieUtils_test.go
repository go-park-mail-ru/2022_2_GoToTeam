package httpCookieUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeHttpCookie(t *testing.T) {
	sessionId := "qwertyzxc"
	cookie := MakeHttpCookie(sessionId)
	if cookie == nil {
		t.Error()
	}

	assert.Equal(t, sessionId, cookie.Value)
}

//func TestExpireHttpCookie(t *testing.T) {
//	sessionId := "qwertyzxc"
//	cookie := MakeHttpCookie(sessionId)
//
//	_, _, currentDay := cookie.Expires.Date()
//	ExpireHttpCookie(cookie)
//	_, _, day := cookie.Expires.Date()
//
//	if currentDay <= day {
//		t.Error()
//	}
//}
