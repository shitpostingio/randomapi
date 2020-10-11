package main

import (
	"os"

	"github.com/shitpostingio/randomapi/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	c              *config.Config
	debug          bool
	allowedOrigins []string

	configFilePath string

	memesCollection *mongo.Collection

	requestedMemes map[string]*os.File
)
