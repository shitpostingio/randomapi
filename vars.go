package main

import (
	"github.com/shitpostingio/randomapi/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	c              *config.Config
	allowedOrigins []string

	configFilePath string

	postCollection *mongo.Collection

	requestedPosts map[string]requestedPost
)
