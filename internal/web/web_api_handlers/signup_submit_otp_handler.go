package web_api_handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/internal/config"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func SignupSubmitOtp(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit submit otp route")
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
			http.Error(w, "We had trouble processing that request.", http.StatusUnprocessableEntity)
			return
		}

		ctx := r.Context()

		email := body.Email

		otp := body.Otp

		validateOtpErr := web_auth.ValidateOtp(ctx, sessionUuid, valkeyClient, email, otp)
		if validateOtpErr != nil {
			http.Error(w, validateOtpErr.Message, validateOtpErr.Code)
			return
		}

		createUserParams := query_handlers.CreateUserParams {
			Email: email,
		}

		createdUser, createUserErr := query_handlers.HandleCreateUser(ctx, pgCoreQueries, pgCore, createUserParams)
		if createUserErr != nil {
			http.Error(w, createUserErr.Message, createUserErr.Code)
			return
		}

		loginUserErr := web_auth.LoginUser(ctx, sessionStore, valkeyClient, r, w, createdUser.UserId)
		if loginUserErr != nil {
			http.Error(w, loginUserErr.Message, loginUserErr.Code)
			return
		}

		http.Redirect(w, r, "/profile", 303)
	}
}
