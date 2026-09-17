package web_page_handlers

import (
	"freecreate/internal/lib/logger"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func EditProfilePageHandler(templates *template.Template)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		type PageData struct{
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn: true,
				LoggedInClass: "logged_in",
				CsrfToken: csrf.TemplateField(r),
			},
		}

		templErr := templates.ExecuteTemplate(w, "edit_profile_page", pageData)
		if templErr != nil {
			logger.Log(templErr)
		}
	}
}