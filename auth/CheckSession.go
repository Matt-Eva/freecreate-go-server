package auth

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func CheckSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	session, _ := sessionStore.Get(r, "user-session")
	
	if session.Values["userId"] == nil {
		return uuid.Nil, errors.New("user not logged in")
	}

	val := session.Values["userId"]

	userId, ok := val.(uuid.UUID)
	if !ok {
		session.Values = make(map[interface{}]interface{})
		session.Options.MaxAge = -1
		return uuid.Nil, errors.New("session does not contain a valid uuid - session destroyed")
	}

	session.Options.MaxAge = 3600 * 12

	err := session.Save(r, w)

	return userId, err
}
