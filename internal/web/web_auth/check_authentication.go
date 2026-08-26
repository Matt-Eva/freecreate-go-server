package web_auth

import (
	"context"
	"freecreate/internal/lib/api_error"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func CheckAuthentication(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) (int64, *api_error.Error) {
	_, userId, _ := GetUser(ctx, sessionStore, valkeyClient, w, r)
	if userId != 0 {
		return userId, nil
	}

	guestSessionUuid, checkGuestSessionErr := CheckGuestSession(sessionStore, w, r)
	if guestSessionUuid == uuid.Nil {
		return 0, checkGuestSessionErr
	}

	return 0, nil
}
