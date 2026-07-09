package web_page_handlers

import (
	web_page_utils "freecreate/web_page_handlers/utils"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func SignupPageHandler(signupTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleSignupPageGet(signupTmpl, w, r)
		case "POST":
			handleSignupPagePost(signupTmpl, w, r)
		default:
			web_page_utils.HandleInvalidWebpageRequestMethod(w, r.Method)
		}
	}
}

func renderSignupPage(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		CsrfToken template.HTML
	}

	pageData := PageData{
		CsrfToken: csrf.TemplateField(r),
	}

	signupTmpl.ExecuteTemplate(w, "signup_page", pageData)
}

func handleSignupPageGet(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	renderSignupPage(signupTmpl, w, r)
}

func handleSignupPagePost(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {

}
