package web_page_handlers

import (
	"freecreate/internal/lib/logger"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func ChapterPageHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      false,
				LoggedInClass: "logged_out",
				CsrfToken:     csrf.TemplateField(r),
			},
		}

		err := templates.ExecuteTemplate(w, "chapter_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
