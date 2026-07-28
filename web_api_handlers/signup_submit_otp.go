package web_api_handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/auth"
	"freecreate/config"
	pg_core_queries "freecreate/db/pg_core/queries"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func SignupSubmitOtp(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit submit otp route")
		_, sessionUuid, getSessionErr := auth.GetGuestSession(sessionStore, w, r)
		if getSessionErr != nil {
			logger.Log(getSessionErr)
			http.Error(w, getSessionErr.Error(), 500)
			return
		}

		type RequestBody struct {
			Email string `json:"email"`
			Otp   string `json:"otp"`
		}

		var body RequestBody

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), 422)
			return
		}

		ctx := r.Context()

		email := body.Email
		emailValidationError := pg_core_validators.ValidateEmail(email)
		if emailValidationError != nil {
			logger.Log(emailValidationError)
			http.Error(w, emailValidationError.Error(), 422)
			return
		}
		fmt.Println("email successfully validated!")

		otp := body.Otp

		validateOtpErr := auth.ValidateOtp(ctx, sessionUuid, valkeyClient, email, otp)
		if validateOtpErr != nil {
			logger.Log(validateOtpErr)
			http.Error(w, validateOtpErr.Error(), 500)
			return
		}
		fmt.Println("otp successfully validated!")

		userId, createUserErr := pg_core_queries.CreateUser(ctx, pgCoreQueries, pgCore, email)
		if createUserErr != nil {
			logger.Log(createUserErr)
			http.Error(w, createUserErr.Error(), 500)
			return
		}
		fmt.Println("user successfully created!")
		fmt.Println(userId)

		loginUserErr := auth.LoginUser(ctx, sessionStore, valkeyClient, r, w, userId)
		if loginUserErr != nil {
			logger.Log(loginUserErr)
			http.Error(w, loginUserErr.Error(), 500)
			return
		}
		fmt.Println("user successfully logged in!")

		
		fmt.Println("redirecting")
		http.Redirect(w, r, "/profile", 303)
	}
}
