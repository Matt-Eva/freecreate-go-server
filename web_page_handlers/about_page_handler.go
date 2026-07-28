package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func AboutPageHandler(aboutTmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isLoggedIn, _ := auth.CheckLogin(sessionStore, w, r)

		type PageData struct {
			Title    string
			LoggedIn bool
		}

		pageData := PageData{
			Title:    "about",
			LoggedIn: isLoggedIn,
		}

		aboutTmpl.ExecuteTemplate(w, "about_page", pageData)
	}
}

func write103Header(w http.ResponseWriter) {
	w.Header().Add("Link", "</static/globals.css>; rel=preload; as=style")
	w.Header().Add("Link", "</static/header_component.css>; rel=preload; as=style")
	w.WriteHeader(103)
}
