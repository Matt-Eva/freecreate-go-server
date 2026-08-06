package web_api_handlers

import (
	"encoding/json"
	"fmt"

	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/logger"
	"freecreate/internal/web/web_auth"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func LoginSubmitOtpHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit login submit otp route")
		_, sessionUuid, getSessionErr := web_auth.GetGuestSession(sessionStore, w, r)
		if getSessionErr != nil {
			http.Error(w, getSessionErr.Message, getSessionErr.Code)
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
			http.Error(w, emailValidationError.Message, emailValidationError.Code)
			return
		}

		otp := body.Otp

		validateOtpErr := web_auth.ValidateOtp(ctx, sessionUuid, valkeyClient, email, otp)
		if validateOtpErr != nil {
			http.Error(w, validateOtpErr.Message, validateOtpErr.Code)
			return
		}

		userId, checkEmailErr := pg_core_queries.GetUserByEmail(ctx, pgCoreQueries, pgCore, email)
		if checkEmailErr != nil {
			http.Error(w, checkEmailErr.Message, checkEmailErr.Code)
			return
		}

		loginUserErr := web_auth.LoginUser(ctx, sessionStore, valkeyClient, r, w, userId)
		if loginUserErr != nil {
			http.Error(w, loginUserErr.Message, loginUserErr.Code)
			return
		}

		http.Redirect(w, r, "/profile", 303)

	}

}
