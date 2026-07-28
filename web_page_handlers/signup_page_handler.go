package web_page_handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func SignupPageHandler(signupTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			CsrfToken template.HTML
		}

		pageData := PageData{
			CsrfToken: csrf.TemplateField(r),
		}

		signupTmpl.ExecuteTemplate(w, "signup_page", pageData)
	}
}

func renderSignupPage(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {

}
