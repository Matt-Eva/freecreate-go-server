package auth

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func GetUser(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	session, _ := GetSession(sessionStore, r)

	if session.Values["userId"] == nil {
		session.Options.MaxAge = -1
		session.Save(r, w)
		return nil, errors.New("user not logged in")
	}

	val := session.Values["userId"]

	_, ok := val.(uuid.UUID)
	if !ok {
		session.Values["userId"] = nil
		session.Options.MaxAge = -1
		session.Save(r, w)
		return nil, errors.New("session does not contain a valid uuid - session destroyed")
	}

	return session, nil
}
