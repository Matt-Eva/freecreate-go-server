package routes

import (
	"freecreate/internal/config"
	"freecreate/internal/web/middleware"
	"freecreate/internal/web/web_api_handlers"
	"freecreate/internal/web/web_page_handlers"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func ConfigureWebRouter(router chi.Router, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgxPools config.PgxPools, pgCoreQueries config.PgCoreQueries, resendClient *resend.Client) {

	router.Group(func(router chi.Router) {
		// ========= Router Configuration ========

		csrfMiddleware := middleware.GenereateCsrfMiddleware()
		router.Use(csrfMiddleware)

		fileServer := http.FileServer(http.Dir("internal/web/static"))
		cachedFileServer := middleware.CacheControlHandler(fileServer)

		router.Handle("/internal/web/static/*", http.StripPrefix("/internal/web/static/", cachedFileServer))

		templates := template.Must(template.ParseGlob("internal/web/templates/*html"))

		// ========= Web Page Handlers =========

		router.Get("/", web_page_handlers.HomePageHandler(templates, sessionStore, valkeyClient))
		router.Get("/browse/{writing_type}", web_page_handlers.HomePageHandler(templates, sessionStore, valkeyClient))

		router.Get("/login", web_page_handlers.LoginPageHandler(sessionStore, valkeyClient, templates))

		router.Get("/signup", web_page_handlers.SignupPageHandler(templates, sessionStore, valkeyClient))

		router.Get("/profile", web_page_handlers.ProfilePageHandler(sessionStore, valkeyClient, templates))

		router.Get("/about", web_page_handlers.AboutPageHandler(templates, sessionStore, valkeyClient))

		router.Get("/donate", web_page_handlers.DonatePageHandler(templates, sessionStore, valkeyClient))

		router.Get("/search", web_page_handlers.SearchPageHandler(templates, sessionStore, valkeyClient))

		router.Get("/my-creators", web_page_handlers.MyCreatorsPageHandler(templates, sessionStore, valkeyClient))

		// ======== JSON Web API Routes =========

		router.Route("/web-api", func(r chi.Router) {

			r.Post("/signup/request-otp", web_api_handlers.SignupRequestOtp(sessionStore, valkeyClient, resendClient, pgCoreQueries, pgxPools.PgCore))

			r.Post("/signup/submit-otp", web_api_handlers.SignupSubmitOtp(sessionStore, valkeyClient, pgCoreQueries, pgxPools.PgCore))

			r.Post("/login/request-otp", web_api_handlers.LoginRequestOtpHandler(sessionStore, valkeyClient, resendClient, pgCoreQueries, pgxPools.PgCore))

			r.Post("/login/submit-otp", web_api_handlers.LoginSubmitOtpHandler(sessionStore, valkeyClient, pgCoreQueries, pgxPools.PgCore))

			r.Delete("/logout", web_api_handlers.LogoutHandler(sessionStore, valkeyClient))

			// r.Post("/creator", web_api_)
		})

	})
}
