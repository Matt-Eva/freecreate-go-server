package auth

import (
	"errors"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func GetGuestSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, uuid.UUID, *api_error.Error) {
	session, err := sessionStore.Get(r, "guest-session")
	if err != nil {
		logger.Log(err)
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: err,
		}
		return nil, uuid.UUID{}, apiErr
	}

	sessionUuid, ok := session.Values["session_uuid"].(uuid.UUID)
	if !ok {
		err := errors.New("Could not convert uuid value - destroying session")
		logger.Log(err)

		destroyErr := DestroyGuestSession(sessionStore, w, r)
		if destroyErr != nil {
			return nil, uuid.UUID{}, destroyErr
		}

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: err,
		}

		return nil, uuid.UUID{}, apiErr
	}

	return session, sessionUuid, nil
}
