package auth

import (
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

func DestroyUserSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) *api_error.Error {
	err := destroySession(session, w, r, "user-session")
	
	return err
}

func DestroyGuestSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) *api_error.Error {
	err := destroySession(session, w, r, "guest-session")

	return err
}

func destroySession(session *sessions.Session, w http.ResponseWriter, r *http.Request, sessionName string) *api_error.Error {
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		logger.Log(err)
		return &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: err,
		}
	}

	return &api_error.Error{}
}
