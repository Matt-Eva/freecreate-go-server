package web_page_handlers

import (
	"html/template"
	"net/http"
)

func AboutPageHandler(aboutTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type PageData struct {
			Title    string
			LoggedIn bool
		}

		pageData := PageData{
			Title:    "about",
			LoggedIn: false,
		}

		aboutTmpl.ExecuteTemplate(w, "about", pageData)
	}
}
