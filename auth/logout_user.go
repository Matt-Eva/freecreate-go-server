package auth

import (
	"context"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LogoutUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) error {
	session, _, getSessionErr := GetSession(sessionStore, w, r)
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		return getSessionErr
	}

	// sessionUuid := session.Values["session_uuid"]

	destroySessionErr := DestroyUserSession(session, w, r)
	if destroySessionErr != nil {
		logger.Log(destroySessionErr)
		return destroySessionErr
	}

	return nil
}
