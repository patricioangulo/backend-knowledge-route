package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func Connect(defaultURI string) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default URI")
	}

	// Override with .env if present
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = defaultURI
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	log.Println("[ATLAS-debug]	Connected to MongoDB Atlas.")
	MongoClient = client
}

func GetCollection(collection string) *mongo.Collection {
	dbName := os.Getenv("MONGO_DB")
	return MongoClient.Database(dbName).Collection(collection)
}
