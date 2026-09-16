package web_page_handlers

import (
	"freecreate/internal/lib/logger"
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

	err := template.ExecuteTemplate(w, "error_page", pageData)
	if err != nil {
		logger.Log(err)
	}
}
