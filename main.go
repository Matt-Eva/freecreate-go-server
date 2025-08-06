package main

import (
	"context"
	"freecreate/config"
	"freecreate/middleware"
	"freecreate/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	gormPGClient := config.ConfigPG()

	mongoClient := config.ConfigMongo()

	valkeyClient := config.ConfigValkey()

	resendClient := config.InitResend()

	router := routes.CreateRouter(gormPGClient, mongoClient, valkeyClient, resendClient)

	var srv *http.Server

	if environment != "PRODUCTION" {
		corsRouter := middleware.DevCorsMiddleware(router)
		srv = &http.Server{
			Addr:         ":8080",
			Handler:      corsRouter,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}
	} else {
		srv = &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}
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

	gormPGDB, err := gormPGClient.DB()
	if err != nil {
		log.Fatalf("could not access gorm pg db: %v", err)
	}
	gormPGDB.Close()
	log.Println("pg db connection shutdown")

	log.Println("main function closing gracefully. Goodbye!")
}
