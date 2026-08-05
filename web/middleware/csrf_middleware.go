package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
)

func GenereateCsrfMiddleware() func(http.Handler) http.Handler {
	environment := os.Getenv("ENVIRONMENT")
	csrfKey := os.Getenv("CSRF_KEY")
	var csrfMiddleware func(http.Handler) http.Handler

	if environment == "DEVELOPMENT" {
		fmt.Println("DEVELOPMENT")
		csrfMiddleware = csrf.Protect([]byte(csrfKey), csrf.Secure(false), csrf.TrustedOrigins([]string{"localhost:8080"}))
	} else {
		csrfMiddleware = csrf.Protect([]byte(csrfKey))
	}
	return csrfMiddleware
}
