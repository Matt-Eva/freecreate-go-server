package main

import (
	"context"
	"fmt"
	"freecreate/middleware"
	"freecreate/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



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

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")

	mongoOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect( mongoOptions)
	if err != nil {
		log.Fatalf("error connecting to mongo: %v", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("error pinging mongo: %v", err)
	}
	log.Println("connection to mongo successful!", mongoClient)

	router := routes.CreateRouter()

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

	if err = srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server failed to shutdown: %v", err)
	}
	log.Println("http server shutdown")

	gormPGDB, err := gormPG.DB()
	if err != nil {
		log.Fatalf("could not access gorm pg db: %v", err)
	}
	gormPGDB.Close()
	log.Println("pg db connection shutdown")

	log.Println("main function closing gracefully. Goodbye!")

}
