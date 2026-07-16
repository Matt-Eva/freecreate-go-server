package auth

import (
	"errors"
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func GetUser(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	session, getSessionErr := GetSession(sessionStore, w, r)
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		return nil, getSessionErr
	}

	if session.Values["userId"] == nil {
		return nil, errors.New("user not logged in")
	}

	val := session.Values["userId"]

	_, ok := val.(uuid.UUID)
	if !ok {
		session.Values["userId"] = nil
		return nil, errors.New("session does not contain a valid uuid - session destroyed")
	}

	return session, nil
}
