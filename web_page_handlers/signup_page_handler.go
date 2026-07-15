package web_page_handlers

import (
	"errors"
	"freecreate/auth"
	"freecreate/config"
	pg_core_queries "freecreate/db/pg_core/queries"
	pg_core_validators "freecreate/db/pg_core/validators"
	email_handler "freecreate/email"
	"freecreate/lib/logger"
	web_page_utils "freecreate/web_page_handlers/utils"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func SignupPageHandler(resendClient *resend.Client, valkeyClient valkey.Client, signupTmpl *template.Template, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userSession, _ := auth.GetUser(sessionStore, w, r)
		if userSession != nil {
			http.Redirect(w, r, "/profile", 303)
			return
		}

		switch r.Method {
		case "GET":
			handleSignupPageGet(signupTmpl, w, r)
		case "POST":
			handleSignupPagePost(resendClient, valkeyClient, signupTmpl, w, r, sessionStore, pgxCore, pgCoreQueries)
		default:
			web_page_utils.HandleInvalidWebpageRequestMethod(w, r.Method)
		}
	}
}

func renderSignupPage(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, otpRequested bool, email string, errors []string) {
	type PageData struct {
		CsrfToken    template.HTML
		OtpRequested bool
		Email        string
		Errors       []string
	}

	pageData := PageData{
		CsrfToken:    csrf.TemplateField(r),
		OtpRequested: otpRequested,
		Email:        email,
		Errors:       errors,
	}

	signupTmpl.ExecuteTemplate(w, "signup_page", pageData)
}

func handleSignupPageGet(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	renderSignupPage(signupTmpl, w, r, false, "", []string{})
}

func handleSignupPagePost(resendClient *resend.Client, valkeyClient valkey.Client, signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) {
	formAction, formActionErr := web_page_utils.GetFormAction(r)
	if formActionErr != nil {
		logger.Log(formActionErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{formActionErr.Error()})
		return
	}

	switch formAction {
	case "request_otp":
		handleRequestOtpFormPost(resendClient, valkeyClient, signupTmpl, w, r, sessionStore, pgxCore, pgCoreQueries)
	case "submit_otp":
		handleSubmitOtpFormPost(sessionStore, w, r)
	default:
		err := errors.New("that is not a correct form action for the signup page")
		logger.Log(err)
		renderSignupPage(signupTmpl, w, r, false, "", []string{err.Error()})
		return
	}
}

func handleRequestOtpFormPost(resendClient *resend.Client, valkeyClient valkey.Client, signupTmpl *template.Template, w http.ResponseWriter, r *http.Request, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) {
	email := r.Form.Get("enter_email")
	emailErr := pg_core_validators.ValidateEmail(email)
	if emailErr != nil {
		logger.Log(emailErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{emailErr.Error()})
		return
	}

	_, getUserErr := pg_core_queries.GetUserByEmail(r.Context(), pgCoreQueries, pgxCore, email)

	if errors.Is(getUserErr, pgx.ErrNoRows) {
		requestOtp(resendClient, valkeyClient, sessionStore, w, r, signupTmpl, email)
	} else if getUserErr != nil {
		errMsg := "Our server had trouble processing that request. Please try again."
		renderSignupPage(signupTmpl, w, r, false, "", []string{errMsg})
	} else {
		errMsg := "That email address is already in use. Please enter a different email address to create a new account, or login with your existing email."
		renderSignupPage(signupTmpl, w, r, false, "", []string{errMsg})
	}
}

func requestOtp(resendClient *resend.Client, valkeyClient valkey.Client, sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request, signupTmpl *template.Template, email string){
	session, getSessionErr := auth.GetSession(sessionStore, w, r)
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{getSessionErr.Error()})
		return
	}

	otp, genOtpErr := auth.GenerateOtp()
	if genOtpErr != nil {
		logger.Log(genOtpErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{genOtpErr.Error()})
		return
	}

	storeOtpErr := auth.StoreOtp(valkeyClient, session, email, otp)
	if storeOtpErr != nil {
		logger.Log(storeOtpErr)
		renderSignupPage(signupTmpl, w, r, false, "", []string{storeOtpErr.Error()})
		return
	}

	sendOtpErr := email_handler.SendOtp(resendClient, email, otp)
	if sendOtpErr != nil {
		logger.Log(sendOtpErr)
		renderSignupPage(signupTmpl,w, r, false, "", []string{sendOtpErr.Error()})
		return
	}

	session.Save(r, w)
	renderSignupPage(signupTmpl, w, r, true, email, []string{})
}

func handleSubmitOtpFormPost(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
}
