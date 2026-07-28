package auth

import (
	"errors"
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func GetGuestSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, uuid.UUID, error) {
	session, err := sessionStore.Get(r, "guest-session")
	if err != nil {
		logger.Log(err)
		return nil, uuid.UUID{}, err
	}

	sessionUuid, ok := session.Values["session_uuid"].(uuid.UUID)
	if !ok {
		err := errors.New("Could not convert uuid value - destroying session")
		logger.Log(err)
		destroyErr := DestroyGuestSession(session, w, r)
		if destroyErr != nil {
			logger.Log(destroyErr)
			return nil, uuid.UUID{}, destroyErr
		}
		return nil, uuid.UUID{}, err
	}

	return session, sessionUuid, nil
}
