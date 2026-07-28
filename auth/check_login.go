package auth

import (
	"fmt"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

func CheckLogin(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (bool, error) {
	session, _, getSessionErr := GetSession(sessionStore, w, r)
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		return false, getSessionErr
	}

	loggedIn, _ := session.Values["logged_in"].(bool)
	if loggedIn != true {
		return false, nil
	}
	fmt.Println("is logged in")
	fmt.Println(loggedIn)

	return loggedIn, nil
}
