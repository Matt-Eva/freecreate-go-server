package web_page_handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(homeTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type PageData struct {
			Title    string
			LoggedIn bool
		}

		pageData := PageData{
			Title:    "home",
			LoggedIn: true,
		}
		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
