package web_page_handlers

import (
	"errors"
	"fmt"
	"freecreate/config"
	pg_core_queries "freecreate/db/pg_core/queries"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/logger"
	web_page_utils "freecreate/web_page_handlers/utils"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SignupPageHandler(signupTmpl *template.Template, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleSignupPageGet(signupTmpl, w, r)
		case "POST":
			handleSignupPagePost(signupTmpl, w, r, sessionStore, pgxCore, pgCoreQueries)
		default:
			web_page_utils.HandleInvalidWebpageRequestMethod(w, r.Method)
		}
	}
}

func renderSignupPage(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, otpRequested bool, email string, errors []string) {
	type PageData struct {
		CsrfToken    template.HTML
		OtpRequested bool
		Email string
		Errors       []string
	}

	pageData := PageData{
		CsrfToken:    csrf.TemplateField(r),
		OtpRequested: otpRequested,
		Email: email,
		Errors:       errors,
	}

	signupTmpl.ExecuteTemplate(w, "signup_page", pageData)
}

func handleSignupPageGet(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	renderSignupPage(signupTmpl, w, r, false, "", []string{})
}

func handleSignupPagePost(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) {
	formAction, formActionErr := web_page_utils.GetFormAction(r)
	if formActionErr != nil {
		logger.Log(formActionErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{formActionErr.Error()})
		return
	}

	switch formAction {
	case "request_otp":
		handleRequestOtpFormPost(signupTmpl, w, r, sessionStore, pgxCore, pgCoreQueries)
	case "submit_otp":
		handleSubmitOtpFormPost(sessionStore, w, r)
	default:
		err := errors.New("that is not a correct form action for the signup page")
		logger.Log(err)
		renderSignupPage(signupTmpl, w, r, false, "", []string{err.Error()})
		return
	}
}

func handleRequestOtpFormPost(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) {
	email := r.Form.Get("enter_email")
	fmt.Println(email)
	emailErr := pg_core_validators.ValidateEmail(email)
	if emailErr != nil {
		logger.Log(emailErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{emailErr.Error()})
		return 
	}

	_, getUserErr := pg_core_queries.GetUserByEmail(r.Context(), pgCoreQueries, pgxCore, email)

	if errors.Is(getUserErr, pgx.ErrNoRows){
		renderSignupPage(signupTmpl, w, r, true, email, []string{})
	} else if getUserErr != nil {
		errMsg := "Our server had trouble processing that request. Please try again."
		renderSignupPage(signupTmpl, w, r, false, "", []string{errMsg})
	} else {
		errMsg := "That email address is already in use. Please enter a different email address to create a new account, or login with your existing email."
		renderSignupPage(signupTmpl, w, r, false, "", []string{errMsg})
	}
	
}

func handleSubmitOtpFormPost(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {}
