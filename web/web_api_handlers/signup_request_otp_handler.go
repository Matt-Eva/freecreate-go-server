package web_api_handlers

import (
	"encoding/json"
	"errors"
	"freecreate/config"
	pg_core_queries "freecreate/db/pg_core/queries"
	pg_core_validators "freecreate/db/pg_core/validators"
	email_handler "freecreate/email"
	"freecreate/lib/logger"
	"freecreate/web/web_auth"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func SignupRequestOtp(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, resendClient *resend.Client, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if user already logged in, redirect to profile page
		ctx := r.Context()

		_, userId, _ := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if userId != 0 {
			http.Redirect(w, r, "/profile", 303)
			return
		}

		type RequestBody struct {
			Email string `json:"email"`
		}

		var body RequestBody

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
			return
		}

		email := body.Email
		emailValidationErr := pg_core_validators.ValidateEmail(email)
		if emailValidationErr != nil {
			http.Error(w, emailValidationErr.Message, http.StatusUnprocessableEntity)
			return
		}

		_, checkEmailInUseErr := pg_core_queries.GetUserByEmail(ctx, pgCoreQueries, pgCore, email)
		if checkEmailInUseErr != nil && !errors.Is(checkEmailInUseErr.Error, pgx.ErrNoRows) {
			http.Error(w, checkEmailInUseErr.Message, checkEmailInUseErr.Code)
			return
		} else if checkEmailInUseErr == nil {
			http.Error(w, "Email address already in use.", http.StatusUnprocessableEntity)
			return
		}

		_, sessionUuid, getSessionErr := web_auth.CreateGuestSesion(sessionStore, w, r)
		if getSessionErr != nil {
			http.Error(w, getSessionErr.Message, getSessionErr.Code)
			return
		}

		otp, genOtpErr := web_auth.GenerateOtp()
		if genOtpErr != nil {
			http.Error(w, genOtpErr.Message, genOtpErr.Code)
			return
		}

		storeOtpErr := web_auth.StoreOtp(ctx, valkeyClient, sessionUuid, email, otp)
		if storeOtpErr != nil {
			http.Error(w, storeOtpErr.Message, storeOtpErr.Code)
			return
		}

		sendEmailErr := email_handler.SendOtp(resendClient, email, otp)
		if sendEmailErr != nil {
			if sendEmailErr.Error() == "[ERROR]: Invalid `to` field. The email address needs to follow the `email@example.com` or `Name <email@example.com>` format." {
				err := errors.New("That is not a valid email address. Please enter a valid email address.")
				http.Error(w, err.Error(), 422)
				return
			}

			http.Error(w, sendEmailErr.Error(), 422)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
