package web_page_handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func MyLibraryPageHandler(templates *template.Template) http.HandlerFunc {
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

		templates.ExecuteTemplate(w, "my_library_page", pageData)
	}
}
