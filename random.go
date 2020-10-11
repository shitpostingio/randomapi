package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/shitpostingio/randomapi/rest/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Random will return a random meme
func random(w http.ResponseWriter, r *http.Request) {

	meme, err := getRandomMeme(r.Host, memesCollection)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(meme)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}

	return
}

func getRandomMeme(host string, collection *mongo.Collection) (client.Response, error) {

	


	return nil, nil
}

func findPost(collection *mongo.Collection) (Post, error) {
	//
	ctx, cancelCtx := context.WithTimeout(context.Background(), opDeadline)
	defer cancelCtx()

	//
	filter := bson.M{"media.fileuniqueid": uniqueID}

	//
	result := collection.FindOne(ctx, filter, options.FindOne())
	if result.Err() != nil {
		return post, result.Err()
	}

	//
	err = result.Decode(&post)
	return post, err
}
