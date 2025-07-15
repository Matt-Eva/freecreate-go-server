package routes

import (
	"encoding/json"
	"fmt"
	"freecreate/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/resend/resend-go/v2"
)

func CreateRouter(resendClient *resend.Client) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/email", handlers.EmailHandler(resendClient)).Methods("POST")

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit route")

		type Response struct {
			Message string `json:"message"`
		}

		response := Response{
			Message: "Hello world!",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET")

	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	return router
}
