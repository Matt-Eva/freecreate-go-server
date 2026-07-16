package web_page_handlers

import (
	"errors"
	"fmt"
	"freecreate/auth"
	"freecreate/lib/logger"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func SignupVerifyOtpPageHandler(signupVerifyOtpTmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, getSessionErr := auth.GetSession(sessionStore, w, r)
		if getSessionErr != nil {
			logger.Log(getSessionErr)
			http.Redirect(w, r, "/signup", 303)
			return
		}

		email := session.Flashes("email")
		session.Save(r, w)
		fmt.Println(email)
		if len(email) == 0 {
			err := errors.New("no email was passed in session flash - redirecting to signup")
			logger.Log(err)
			http.Redirect(w, r, "/signup", 303)
			return
		}
		// signupVerifyOtpTmpl
		signupVerifyOtpTmpl.ExecuteTemplate(w, "signup_verify_otp_page", map[string]string{})
	}
}
