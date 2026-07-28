package web_api_handlers

import (
	"freecreate/auth"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)



func LogoutHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		_, _, getUserErr := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if getUserErr != nil {
			logger.Log(getUserErr)
			http.Error(w, getUserErr.Error(), 500)
			return
		}

		
		
		logoutErr := auth.LogoutUser(ctx, sessionStore, valkeyClient, w, r)
		if logoutErr != nil {
			logger.Log(logoutErr)
			http.Error(w, logoutErr.Error(), 500)
			return
		}

		http.Redirect(w, r, "/", 303)
	}
}
