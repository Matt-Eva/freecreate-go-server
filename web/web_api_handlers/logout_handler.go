package web_api_handlers

import (
	"freecreate/web/web_auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LogoutHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, _, getUserErr := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if getUserErr != nil {

			http.Error(w, getUserErr.Message, getUserErr.Code)
			return
		}

		logoutErr := web_auth.LogoutUser(ctx, sessionStore, valkeyClient, w, r)
		if logoutErr != nil {

			http.Error(w, logoutErr.Message, logoutErr.Code)
			return
		}

		http.Redirect(w, r, "/", 303)
	}
}
