package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/shitpostingio/randomapi/backstore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	setCLIParams()

	if !debug {
		allowedOrigins = append(allowedOrigins, "*")
	} else {
		allowedOrigins = append(allowedOrigins, "random.shitposting.io")
	}

	if err := envSetup(); err != nil {
		log.Fatal(fmt.Errorf("cannot detect env: %w", err))
	}

	client, err := mongo.Connect(context.Background(), c.Mongo.MongoDBConnectionOptions())
	if err != nil {
		log.Fatal("Unable to connect to document store:", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Unable to ping document store:", err)
	}

	collection := client.Database(c.Mongo.DatabaseName).Collection(c.Mongo.CollectionName)

	bs, err = backstore.NewBackstore(c, collection)
	if err != nil {
		err = fmt.Errorf("cannot build backstore: %w", err)
		log.Fatal(err)
	}

	feedErrors = make(chan error)
}

func setCLIParams() {
	flag.StringVar(&path, "config", "./config.toml", "configuration file path")
	flag.BoolVar(&debug, "dev", false, "developer debug mode")
	flag.Parse()
}
