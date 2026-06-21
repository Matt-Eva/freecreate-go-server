package routes

import (
	"fmt"
	"freecreate/middleware"
	"freecreate/web_api_handlers"
	"freecreate/web_page_handlers"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

func CreateRouter(sessionStore *sessions.CookieStore,pgxMainDb *pgxpool.Pool, pgxContentDbOne *pgxpool.Pool, gormPGClient *gorm.DB, mongoClient *mongo.Client, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {
	router := chi.NewRouter()

	environment := os.Getenv("ENVIRONMENT")

	csrfKey := os.Getenv("CSRF_KEY")
	var csrfMiddleware func(http.Handler) http.Handler

	if environment == "DEVELOPMENT" {
		fmt.Println("DEVELOPMENT")
		csrfMiddleware = csrf.Protect([]byte(csrfKey), csrf.Secure(false), csrf.TrustedOrigins([]string{"localhost:8080"}))
	} else {
		csrfMiddleware = csrf.Protect([]byte(csrfKey))
	}

	router.Use(csrfMiddleware)

	fileServer := http.FileServer(http.Dir("static"))
	cachedFileServer := middleware.CacheControlHandler(fileServer)

	router.Handle("/static/*", http.StripPrefix("/static/", cachedFileServer))

	router.Get("/get-csrf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
	})

	templates := template.Must(template.ParseGlob("templates/*html"))

	router.Get("/", web_page_handlers.HomePageHandler(templates))

	router.Get("/test", web_page_handlers.TestPageHandler(templates))
	router.Post("/test", web_page_handlers.TestPageHandler(templates))

	router.Get("/about", web_page_handlers.AboutPageHandler(templates))

	router.Get("/login", web_page_handlers.LoginPageHandler(templates))

	router.Get("/profile", web_page_handlers.ProfilePageHandler(templates))

	router.Get("/donate", web_page_handlers.DonatePageHandler(templates))

	// ======== JSON Web API Routes =========
	router.Route("/web-api", func(r chi.Router) {

		r.Post("/test", web_api_handlers.TestHandler())

		r.Post("/login", web_api_handlers.LoginHandler(sessionStore, gormPGClient))

		r.Post("/signup", web_api_handlers.SignupHandler(sessionStore, gormPGClient))

		r.Delete("/logout", web_api_handlers.LogoutHandler(sessionStore, gormPGClient))

		r.Get("/reauth", web_api_handlers.ReAuthHandler(sessionStore, gormPGClient))

		r.Delete("/delete-account", web_api_handlers.DeleteAccountHandler(sessionStore, gormPGClient))

		r.Post("/creator", web_api_handlers.CreateCreatorHandler(sessionStore, gormPGClient))

		r.Delete("/creator/{creatorId}", web_api_handlers.DeleteCreatorHandler(sessionStore, gormPGClient))

		r.Get("/user-creators", web_api_handlers.GetUserCreatorHandlers(sessionStore, gormPGClient))

		r.Post("/writing", web_api_handlers.CreateWritingHandler(sessionStore, gormPGClient))

		r.Patch("/writing", web_api_handlers.UpdateWritingHandler(sessionStore, gormPGClient))

		r.Get("/my-writing", web_api_handlers.GetMyWritingHandler(sessionStore, gormPGClient))

		r.Get("/edit-writing/{writingUUID}", web_api_handlers.GetEditWritingHandler(sessionStore, gormPGClient))

		// router.Post("/createOTP", handlers.CreateOTPHandler(resendClient, valkeyClient))

		// router.Post("/email", handlers.EmailHandler(resendClient))

	})

	return router
}
