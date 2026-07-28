package web_api_handlers

import (
	"encoding/json"
	"errors"
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

func LoginRequestOtpHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if user already logged in, redirect to profile page
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
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
			logger.Log(emailValidationErr)
			http.Error(w, emailValidationErr.Error(), http.StatusUnprocessableEntity)
			return
		}

		userId, checkEmailErr := pg_core_queries.GetUserByEmail(ctx, pgCoreQueries, pgCore, email)
		if checkEmailErr != nil {
			logger.Log(checkEmailErr)
			err := errors.New("we had trouble finding an account with that address")
			http.Error(w, err.Error(), 404)
			return
		}

	}
}
