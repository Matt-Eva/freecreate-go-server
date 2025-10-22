package routes

import (
	"encoding/json"
	"fmt"
	"freecreate/handlers"
	"freecreate/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

func CreateRouter(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB, mongoClient *mongo.Client, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/createOTP", handlers.CreateOTPHandler(resendClient, valkeyClient))

	router.Post("/email", handlers.EmailHandler(resendClient))

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

	router.Post("/login", handlers.LoginHandler(sessionStore, gormPGClient))

	router.Post("/signup", func(w http.ResponseWriter, r *http.Request){

	})

	router.Delete("/logout", func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessionStore.Get(r, "user-session")
		session.Values = make(map[interface{}]interface{})
		session.Options.MaxAge = -1

		err := session.Save(r, w)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
