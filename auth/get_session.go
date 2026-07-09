package auth

import (
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

func GetSession(sessionStore *sessions.CookieStore, r *http.Request) (*sessions.Session, error) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		logger.Log(err)
		return nil, err
	}
	session.Options.MaxAge = 3600 * 12
	return session, nil
}
