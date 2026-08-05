package auth

import (
	"context"
	"freecreate/lib/api_error"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LogoutUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) *api_error.Error {
	
	destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
	if destroySessionErr != nil {
		return destroySessionErr
	}

	return nil
}
