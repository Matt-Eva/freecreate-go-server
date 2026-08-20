package web_page_handlers

import (
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func AboutPageHandler(aboutTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)

		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		} 

		type PageData struct {
			Title         string
			LoggedIn      bool
			LoggedInClass string
		}

		pageData := PageData{
			Title:         "about",
			LoggedIn:      loggedIn,
			LoggedInClass: loggedInClass,
		}

		aboutTmpl.ExecuteTemplate(w, "about_page", pageData)
	}
}

func write103Header(w http.ResponseWriter) {
	w.Header().Add("Link", "</static/globals.css>; rel=preload; as=style")
	w.Header().Add("Link", "</static/header_component.css>; rel=preload; as=style")
	w.WriteHeader(103)
}
