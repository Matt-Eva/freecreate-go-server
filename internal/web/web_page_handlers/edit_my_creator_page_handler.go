package web_page_handlers

import (
	"freecreate/internal/lib/logger"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func EditMyCreatorPageHandler(template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      true,
				LoggedInClass: "logged_in",
				CsrfToken:     csrf.TemplateField(r),
			},
		}

		err := template.ExecuteTemplate(w, "edit_my_creator_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
