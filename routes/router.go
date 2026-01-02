package routes

import (
	"freecreate/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

func CreateRouter(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB, mongoClient *mongo.Client, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/login", handlers.LoginHandler(sessionStore, gormPGClient))

	router.Post("/signup", handlers.SignupHandler(sessionStore, gormPGClient))

	router.Delete("/logout", handlers.LogoutHandler(sessionStore, gormPGClient))

	router.Get("/reauth", handlers.ReAuthHandler(sessionStore, gormPGClient))

	router.Delete("/delete-account", handlers.DeleteAccountHandler(sessionStore, gormPGClient))
	
	router.Post("/creator", handlers.CreateCreatorHandler(sessionStore, gormPGClient))
	
	router.Delete("/creator/{creatorId}", handlers.DeleteCreatorHandler(sessionStore, gormPGClient))
	
	router.Get("/user-creators", handlers.GetUserCreatorHandlers(sessionStore, gormPGClient))
	
	router.Post("/writing", handlers.CreateWritingHandler(sessionStore, gormPGClient))
	
	router.Get("/my-writing", handlers.GetMyWritingHandler(sessionStore, gormPGClient))
	
	router.Get("/edit-writing/{writingUUID}", handlers.GetEditWritingHandler(sessionStore, gormPGClient))
	
	// router.Post("/createOTP", handlers.CreateOTPHandler(resendClient, valkeyClient))

	// router.Post("/email", handlers.EmailHandler(resendClient))

	// router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("hit route hello")

	// 	type Response struct {
	// 		Message string `json:"message"`
	// 	}

	// 	response := Response{
	// 		Message: "Hello world!",
	// 	}

	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(response)
	// })

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
