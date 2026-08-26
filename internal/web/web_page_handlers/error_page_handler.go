package web_page_handlers

import (
	"html/template"
	"net/http"
)

func ErrorPageHandler(template *template.Template, w http.ResponseWriter, message string, loggedIn bool, loggedInClass string) {
	type PageData struct {
		UniversalPageData
	}

	pageData := PageData{
		UniversalPageData: UniversalPageData{
			LoggedIn:      loggedIn,
			LoggedInClass: loggedInClass,
		},
	}

	template.ExecuteTemplate(w, "error_page", pageData)
}
