package routes

import (
	"freecreate/handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

func CreateRouter(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB, mongoClient *mongo.Client, valkeyClient valkey.Client, resendClient *resend.Client) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/createOTP", handlers.CreateOTPHandler(resendClient, valkeyClient)).Methods("POST")

	router.HandleFunc("/email", handlers.EmailHandler(resendClient)).Methods("POST")

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
