package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func devCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := os.Getenv("CLIENT_ORIGIN")

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Contol-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	environment := os.Getenv("ENVIRONMENT")

	router := mux.NewRouter()

	if environment != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		router = devCorsMiddleware(router)
	}

	

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

	// port := os.Getenv("PORT")
	http.ListenAndServe(":8080", router)
}
