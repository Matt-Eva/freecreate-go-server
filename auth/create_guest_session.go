package auth

import (
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func CreateGuestSesion(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, uuid.UUID, error) {
	session, err := sessionStore.Get(r, "guest-session")
	if err != nil {
		logger.Log(err)
		return nil, uuid.UUID{}, err
	}

	sessionUuid, err := uuid.NewRandom()
	if err != nil {
		logger.Log(err)
		return nil, uuid.UUID{}, err
	}

	session.Values["session_uuid"] = sessionUuid
	session.Options.MaxAge = 3600
	session.Save(r, w)

	return session, sessionUuid, nil
}
