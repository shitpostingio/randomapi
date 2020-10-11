package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

	"github.com/shitpostingio/randomapi/rest/client"
	"go.mongodb.org/mongo-driver/mongo"
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

	post, err := findPost(collection)
	if err != nil {
		return client.Response{}, err
	}

	postID := uuid.New().String()

	randompost := client.Response{
		ID: postID,
		Meme: client.Data{
			URL:      fmt.Sprintf("%s/storage/%s", c.Endpoint, post.Media.FileID),
			Filename: "",
			Type:     post.Media.Type,
			Date:     post.PostedAt,
		},
	}

	return randompost, nil
}

func findPost(collection *mongo.Collection) (Post, error) {
	var post Post

	return post, nil
}
