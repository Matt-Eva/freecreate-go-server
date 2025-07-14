package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	if environment != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	host, user, pwd, db, ssl, port := os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"), os.Getenv("PG_SSL"), os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=%s", host, user, pwd, db, ssl, port)

	gormPG, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting gorm to postgres")
	} else {
		fmt.Println("gorm connect to pg successful!", gormPG)
	}

	router := mux.NewRouter()

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

	if environment != "PRODUCTION" {
		corsRouter := devCorsMiddleware(router)
		http.ListenAndServe(":8080", corsRouter)
	} else {
		http.ListenAndServe(":8080", router)
	}

}
