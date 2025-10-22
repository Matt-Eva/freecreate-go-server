package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func CreateSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request, id uint) error {
	session, _ := sessionStore.Get(r, "user-session")
	session.Values["userId"] = id
	err := session.Save(r, w)
	return err
}
