package web_page_handlers

import (
	"html/template"
	"net/http"
)

func LoginVerifyOtpPageHandler(loginVerifyOtpTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginVerifyOtpTmpl.ExecuteTemplate(w, "login_verify_otp_page", map[string]string{})
	}
}
