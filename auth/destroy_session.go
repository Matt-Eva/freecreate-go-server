package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func DestroyUserSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) error {
	err := destroySession(session, w, r, "user-session")
	return err
}

func DestroyGuestSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) error {
	err := destroySession(session, w, r, "guest-session")
	return err
}

func destroySession(session *sessions.Session, w http.ResponseWriter, r *http.Request, sessionName string) error {
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	return err
}
