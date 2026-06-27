package main

import (
	"context"
	"encoding/base64"
	"encoding/gob"
	"freecreate/config"
	"freecreate/logger"
	"freecreate/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		environment = "DEVELOPMENT"
	}

	gob.Register(uuid.UUID{})

	sessionAuthKey, err := base64.StdEncoding.DecodeString(os.Getenv("SESSION_AUTH_KEY"))
	if err != nil {
		logger.Log(err)
		log.Fatal(err.Error())
		return
	}

	sessionEncryptionKey, err := base64.StdEncoding.DecodeString(os.Getenv("SESSION_ENCRYPTION_KEY"))
	if err != nil {
		logger.Log(err)
		log.Fatal(err.Error())
		return
	}

	sessionStore := sessions.NewCookieStore(sessionAuthKey, sessionEncryptionKey)

	if environment == "PRODUCTION" {
		sessionStore.Options = &sessions.Options{
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
		}
	}

	ctx := context.Background()

	pgxPools, pgxErr := config.ConfigPgx(ctx, environment)
	if pgxErr != nil {
		logger.Log(pgxErr)
		log.Fatal(pgxErr)
		return
	}

	pgCoreQueries, pgCoreQueryError := config.ConfigPgCoreQueries()
	if pgCoreQueryError != nil {
		logger.Log(pgCoreQueryError)
		return
	}

	valkeyClient := config.ConfigValkey()

	resendClient := config.InitResend()

	router := routes.CreateRouter(sessionStore, pgxPools, pgCoreQueries, valkeyClient, resendClient)

	var srv = &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
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

	// gormPGDB, err := gormPGClient.DB()
	// if err != nil {
	// 	log.Fatalf("could not access gorm pg db: %v", err)
	// }
	// gormPGDB.Close()
	// log.Println("pg db connection shutdown")

	log.Println("main function closing gracefully. Goodbye!")
}
