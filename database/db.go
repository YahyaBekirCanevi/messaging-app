package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MessageCollection *mongo.Collection

// ConnectDB initializes the MongoDB client
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// MongoDB connection string
	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	MongoClient = client
	MessageCollection = client.Database("messaging").Collection("messages")

	log.Println("Connected to MongoDB!")
}

// DisconnectDB disconnects the MongoDB client
func DisconnectDB() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect MongoDB:", err)
		} else {
			log.Println("Disconnected from MongoDB.")
		}
	}
}
