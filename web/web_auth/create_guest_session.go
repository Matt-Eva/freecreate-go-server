package web_auth

import (
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func CreateGuestSesion(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) (*sessions.Session, uuid.UUID, *api_error.Error) {
	session, err := sessionStore.Get(r, "guest-session")
	if err != nil {
		logger.Log(err)
		apiErr := api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}
		return nil, uuid.UUID{}, &apiErr
	}

	sessionUuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		logger.Log(uuidErr)
		apiErr := api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}
		return nil, uuid.UUID{}, &apiErr
	}

	session.Values["session_uuid"] = sessionUuid
	session.Options.MaxAge = 3600
	saveErr := session.Save(r, w)
	if saveErr != nil {
		logger.Log(saveErr)
		apiErr := api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}
		return nil, uuid.UUID{}, &apiErr
	}

	return session, sessionUuid, nil
}
