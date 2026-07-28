package main

import (
	"context"
	"encoding/gob"
	"freecreate/config"
	"freecreate/lib/logger"

	// "freecreate/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/joho/godotenv"
)

func initialize(environment string) (*chi.Mux, error) {
	gob.Register(uuid.UUID{})

	ctx := context.Background()

	sessionStore, sessionErr := config.ConfigSessionStore(environment)
	if sessionErr != nil {
		logger.Log(sessionErr)
		return nil, sessionErr
	}

	pgxPools, pgxErr := config.ConfigPgx(ctx, environment)
	if pgxErr != nil {
		logger.Log(pgxErr)
		log.Fatal(pgxErr)
		return nil, pgxErr
	}

	pgCoreQueries, pgCoreQueryError := config.ConfigPgCoreQueries()
	if pgCoreQueryError != nil {
		logger.Log(pgCoreQueryError)
		return nil, pgCoreQueryError
	}

	valkeyClient := config.ConfigValkey()

	resendClient := config.InitResend()

	router := CreateRouter(sessionStore, pgxPools, pgCoreQueries, valkeyClient, resendClient)

	return router, nil
}

func main() {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		environment = "DEVELOPMENT"
	}

	router, err := initialize(environment)
	if err != nil {
		logger.Log(err)
		return
	}

	var srv = &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		// if environment == "DEVELOPMENT" {
		// 	if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
		// 		log.Fatalf("Server failed: %v", err)
		// 	}
		// } else {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
		// }
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutdown signal received, gracefully shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server failed to shutdown: %v", err)
	}
	log.Println("http server shutdown")

	log.Println("main function closing gracefully. Goodbye!")
}
