package web_page_handlers

import (
	"errors"
	"freecreate/lib/logger"
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

func renderSignupPage(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, otpRequested bool, errors []string) {
	type PageData struct {
		CsrfToken    template.HTML
		OtpRequested bool
		Errors       []string
	}

	pageData := PageData{
		CsrfToken:    csrf.TemplateField(r),
		OtpRequested: otpRequested,
		Errors:       errors,
	}

	signupTmpl.ExecuteTemplate(w, "signup_page", pageData)
}

func handleSignupPageGet(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	renderSignupPage(signupTmpl, w, r, false, []string{})
}

func handleSignupPagePost(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	formAction, formActionErr := web_page_utils.GetFormAction(r)
	if formActionErr != nil {
		logger.Log(formActionErr)
		renderSignupPage(signupTmpl, w, r, false, []string{formActionErr.Error()})
		return
	}

	switch formAction {
	case "request_otp":
		handleRequestOtpFormPost(signupTmpl, w, r)
	case "submit_otp":
		handleSubmitOtpFormPost()
	default:
		err := errors.New("that is not a correct form action for the signup page")
		logger.Log(err)
		renderSignupPage(signupTmpl, w, r, false, []string{err.Error()})
		return
	}
}

func handleRequestOtpFormPost(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	renderSignupPage(signupTmpl, w, r, true, []string{})
}

func handleSubmitOtpFormPost() {}
