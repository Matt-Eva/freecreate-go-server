package web_page_handlers

import (
	"html/template"
	"net/http"
)

func ProfilePageHandler(profileTmpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		type PageData struct {
			LoggedIn bool
		}

		pageData := PageData{
			LoggedIn: true,
		}

		profileTmpl.ExecuteTemplate(w, "profile", pageData)
	}
}
