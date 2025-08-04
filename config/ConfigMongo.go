package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)
func ConfigMongo() *mongo.Client{
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")

	mongoOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(mongoOptions)
	if err != nil {
		log.Fatalf("error connecting to mongo: %v", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("error pinging mongo: %v", err)
	}
	log.Println("connection to mongo successful!")

	return mongoClient
}