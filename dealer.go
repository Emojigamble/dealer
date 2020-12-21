package main

import (
	"context"
	"fmt"
	"github.com/Emojigamble/utility/logger"
	"github.com/Emojigamble/utility/setup"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func main() {
	hostname, _ := os.Hostname()

	// define logger
	log := logger.EmojigambleLogger{
		LogOrigin:       fmt.Sprintf("dealer->%s", hostname),
		ActiveLogLevels: logger.AllLogLevels(),
		LogToDatabase:   false,
	}
	log.Log("Starting dealer...", logger.BaseLogLevel)

	// setup firebase auth client
	_, err := setup.FirebaseAuthClient(context.Background(), "emojigamble-key.json")
	if err != nil {
		log.Log(fmt.Sprint(err), logger.ErrorLogLevel)
		panic(err)
	}

	// connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))
	if err != nil {
		log.Log(fmt.Sprint(err), logger.ErrorLogLevel)
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// TODO: register REST-Endpoints for users
	// TODO: register gRPC service
}
