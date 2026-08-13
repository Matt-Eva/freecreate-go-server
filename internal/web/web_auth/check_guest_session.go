package web_auth

import (
	"freecreate/internal/lib/api_error"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func CheckGuestSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request)(uuid.UUID, *api_error.Error){
	_, sessionUuid, _ := GetGuestSession(sessionStore, w, r)
	if sessionUuid != uuid.Nil {
		return sessionUuid, nil
	}

	_, sessionUuid, createGuestSessionErr := CreateGuestSesion(sessionStore, w, r)
	if createGuestSessionErr != nil {
		return uuid.Nil, createGuestSessionErr
	}

	return sessionUuid, nil
}