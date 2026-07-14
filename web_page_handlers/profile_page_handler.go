package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func ProfilePageHandler(sessionStore *sessions.CookieStore,profileTmpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, getUserErr := auth.GetUser(sessionStore, w, r)
		if getUserErr != nil {
			http.Redirect(w, r, "/login", 303)
			return
		}
		type PageData struct {
			LoggedIn bool
		}

		pageData := PageData{
			LoggedIn: true,
		}

		profileTmpl.ExecuteTemplate(w, "profile", pageData)
	}
}
