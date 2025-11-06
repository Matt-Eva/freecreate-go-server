package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func DestroySession(sessionStore sessions.CookieStore, w http.ResponseWriter, r *http.Request)error{
	session, _ := sessionStore.Get(r, "user-session")
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	return err
}