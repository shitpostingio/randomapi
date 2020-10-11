package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	setCLIParams()

	if err := envSetup(); err != nil {
		log.Fatal(fmt.Errorf("cannot detect env: %w", err))
	}

	mongoClient, err := mongo.Connect(context.Background(), c.MongoMemes.MongoDBConnectionOptions())
	if err != nil {
		log.Fatal("Unable to connect to document store:", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Unable to ping document store:", err)
	}

	errChan = make(chan error)
	requestedPosts = make(map[string]requestedPost)
}

func setCLIParams() {
	flag.StringVar(&configFilePath, "config", "./config.toml", "configuration file path")
	flag.BoolVar(&debug, "dev", false, "developer debug mode")
	flag.Parse()
}
