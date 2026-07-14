package auth

import (
	"errors"
	"fmt"
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func GetSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		logger.Log(err)
		return nil, err
	}

	if session.Values["session_uuid"] == nil {
		fmt.Println("nil session_uuid")
		sessionUuid, err := uuid.NewRandom()
		if err != nil {
			logger.Log(err)
			return nil, err
		}
		session.Values["session_uuid"] = sessionUuid
	}

	val := session.Values["session_uuid"]
	_, ok := val.(uuid.UUID)
	if !ok {
		err := errors.New("Could not convert uuid value - destroying session")
		logger.Log(err)
		destroyErr := DestroySession(sessionStore, w, r)
		if destroyErr != nil {
			logger.Log(destroyErr)
			return nil, destroyErr
		}
		return nil, err
	}

	session.Options.MaxAge = 3600 * 12
	return session, nil
}
