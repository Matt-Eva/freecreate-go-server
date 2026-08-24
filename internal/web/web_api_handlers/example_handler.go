package web_api_handlers

import (
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func ExampleHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		_, userId, _ := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0{
			http.Error(w, "There was an issue with your login session. Please logout and try logging in again.", 401)
			return
		}
	}
}