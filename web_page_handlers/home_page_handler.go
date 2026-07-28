package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func HomePageHandler(homeTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		}

		writingType := chi.URLParam(r, "writing_type")

		type PageData struct {
			LoggedIn bool
			LoggedInClass string
			WritingType string
		}

		pageData := PageData{
			LoggedIn: loggedIn,
			LoggedInClass: loggedInClass,
			WritingType: writingType,
		}
		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
