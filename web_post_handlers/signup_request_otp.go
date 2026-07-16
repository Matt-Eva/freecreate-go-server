package web_post_handlers

import (
	"freecreate/config"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func SignupRequestOtp(resendClient *resend.Client, valkeyClient valkey.Client, sessionStore *sessions.CookieStore, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Form.Get("enter_email")
		emailErr := pg_core_validators.ValidateEmail(email)
		if emailErr != nil {
			logger.Log(emailErr)
			http.Redirect(w, r, "/signup", 303)
			return
		}
		
		http.Redirect(w, r, "/signup/verify-otp", 303)
	}
}
