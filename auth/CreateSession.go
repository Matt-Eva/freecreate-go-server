package auth

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func CreateSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request, userId uuid.UUID) error {
	session, _ := sessionStore.Get(r, "user-session")
	session.Values["userId"] = userId
	err := session.Save(r, w)
	return err
}
