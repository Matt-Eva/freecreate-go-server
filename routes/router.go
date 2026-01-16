package routes

import (
	"encoding/json"
	"fmt"
	"freecreate/handlers"
	"freecreate/logger"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/rs/cors"
	"github.com/valkey-io/valkey-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

func CreateRouter(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB, mongoClient *mongo.Client, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {
	router := chi.NewRouter()

	var csrfMiddleware func(http.Handler) http.Handler
	var corsMiddleware *cors.Cors

	environment := os.Getenv("ENVIRONMENT")
	csrfKey := os.Getenv("CSRF_KEY")

	if environment == "PRODUCTION" {
		csrfMiddleware = csrf.Protect([]byte(csrfKey))
	} else if environment == "DEVELOPMENT" {
		clientOrigin := os.Getenv("CLIENT_ORIGIN")
		
		corsMiddleware = cors.New(cors.Options{
			AllowedOrigins:   []string{clientOrigin},
			AllowedHeaders:   []string{"Content-Type", "X-CSRF-Token"},
			AllowedMethods:   []string{"POST", "PATCH", "GET", "DELETE", "OPTIONS"},
			AllowCredentials: true,
			MaxAge:           86400,
		})
		router.Use(corsMiddleware.Handler)
		csrfMiddleware = csrf.Protect([]byte(csrfKey), csrf.TrustedOrigins([]string{clientOrigin, "localhost:3000", "localhost:8080"}), csrf.Secure(false), csrf.Path("/"))
	} else {
		log.Fatal("bad environment variable load")
	}

	router.Use(csrfMiddleware)

	router.Get("/get-csrf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	router.Post("/login", handlers.LoginHandler(sessionStore, gormPGClient))

	router.Post("/signup", handlers.SignupHandler(sessionStore, gormPGClient))

	router.Delete("/logout", handlers.LogoutHandler(sessionStore, gormPGClient))

	router.Get("/reauth", handlers.ReAuthHandler(sessionStore, gormPGClient))

	router.Delete("/delete-account", handlers.DeleteAccountHandler(sessionStore, gormPGClient))

	router.Post("/creator", handlers.CreateCreatorHandler(sessionStore, gormPGClient))

	router.Delete("/creator/{creatorId}", handlers.DeleteCreatorHandler(sessionStore, gormPGClient))

	router.Get("/user-creators", handlers.GetUserCreatorHandlers(sessionStore, gormPGClient))

	router.Post("/writing", handlers.CreateWritingHandler(sessionStore, gormPGClient))

	router.Patch("/writing", handlers.UpdateWritingHandler(sessionStore, gormPGClient))

	router.Get("/my-writing", handlers.GetMyWritingHandler(sessionStore, gormPGClient))

	router.Get("/edit-writing/{writingUUID}", handlers.GetEditWritingHandler(sessionStore, gormPGClient))

	// router.Post("/createOTP", handlers.CreateOTPHandler(resendClient, valkeyClient))

	// router.Post("/email", handlers.EmailHandler(resendClient))

	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit route hello")

		type Response struct {
			Message string `json:"message"`
		}

		response := Response{
			Message: "Hello world!",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	router.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		type Body struct {
			Message string `json:"message"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
			return
		}

		fmt.Println(body.Message)

		type Response struct {
			Message string `json:"message"`
		}

		response := Response{
			Message: body.Message,
		}

		res, mErr := json.Marshal(response)
		if mErr != nil {
			logger.Log(mErr)
			http.Error(w, mErr.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})

	return router
}

// router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("hit route")

// 		type Response struct {
// 			Message string `json:"message"`
// 		}

// 		response := Response{
// 			Message: "Hello world!",
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(response)
// 	}).Methods("GET")

// 	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

// 	}).Methods("GET")

// 	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

// 	}).Methods("POST")

// 	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

// 	}).Methods("DELETE")
