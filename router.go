package main

import (
	"freecreate/config"
	"freecreate/middleware"
	"freecreate/web_api_handlers"
	"freecreate/web_page_handlers"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func CreateRouter(sessionStore *sessions.CookieStore, pgxPools config.PgxPools, pgCoreQueries config.PgCoreQueries, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {

	// ========= Router Configuration ========

	router := chi.NewRouter()

	csrfMiddleware := middleware.GenereateCsrfMiddleware()
	router.Use(csrfMiddleware)

	fileServer := http.FileServer(http.Dir("static"))
	cachedFileServer := middleware.CacheControlHandler(fileServer)

	router.Handle("/static/*", http.StripPrefix("/static/", cachedFileServer))

	templates := template.Must(template.ParseGlob("templates/*html"))

	// ========= Web Page Handlers =========

	router.Get("/", web_page_handlers.HomePageHandler(templates, sessionStore, valkeyClient))
	router.Get("/browse/{writing_type}", web_page_handlers.HomePageHandler(templates, sessionStore, valkeyClient))

	router.Get("/login", web_page_handlers.LoginPageHandler(sessionStore, valkeyClient, templates))

	router.Get("/signup", web_page_handlers.SignupPageHandler(templates, sessionStore, valkeyClient))

	router.Get("/profile", web_page_handlers.ProfilePageHandler(sessionStore, valkeyClient, templates))

	router.Get("/about", web_page_handlers.AboutPageHandler(templates, sessionStore, valkeyClient))

	router.Get("/donate", web_page_handlers.DonatePageHandler(templates, sessionStore, valkeyClient))

	router.Get("/search", web_page_handlers.SearchPageHandler(templates, sessionStore, valkeyClient))

	// ======== JSON Web API Routes =========

	router.Route("/web-api", func(r chi.Router) {

		r.Post("/login/request-otp", web_api_handlers.LoginRequestOtpHandler(sessionStore, valkeyClient, resendClient, pgCoreQueries, pgxPools.PgCore))

		r.Post("/login/submit-otp", web_api_handlers.LoginSubmitOtpHandler(sessionStore, valkeyClient, pgCoreQueries, pgxPools.PgCore))

		r.Post("/signup/request-otp", web_api_handlers.SignupRequestOtp(sessionStore, valkeyClient, resendClient, pgCoreQueries, pgxPools.PgCore))

		r.Post("/signup/submit-otp", web_api_handlers.SignupSubmitOtp(sessionStore, valkeyClient, pgCoreQueries, pgxPools.PgCore))

		r.Delete("/logout", web_api_handlers.LogoutHandler(sessionStore, valkeyClient))
	})

	return router
}
