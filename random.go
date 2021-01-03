package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shitpostingio/randomapi/backstore"

	"github.com/google/uuid"
	"github.com/shitpostingio/randomapi/rest/client"
	"go.mongodb.org/mongo-driver/mongo"
)

// Random will return a random post
func random(w http.ResponseWriter, r *http.Request) {

	post, err := getRandomPost(postCollection)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}

	return
}

func getRandomPost(collection *mongo.Collection) (client.Response, error) {

	post, err := backstore.FindRandomPost(collection)
	if err != nil {
		return client.Response{}, err
	}

	postID := uuid.New().String()

	var ext string
	if post.Media.Type == "photo" {
		ext = "jpg"
	} else {
		ext = "mp4"
	}

	filename := fmt.Sprintf("%s.%s", post.Media.FileID, ext)

	randompost := client.Response{
		ID: postID,
		Post: client.Data{
			URL:      fmt.Sprintf("%s/storage/%s", c.Endpoint, filename),
			Filename: filename,
			Type:     post.Media.Type,
			Date:     post.PostedAt,
		},
	}

	requestedPosts[filename] = requestedPost{
		path:        fmt.Sprintf("%s/%s", c.PostFolder, filename),
		mediatype:   post.Media.Type,
		requestdate: time.Now(),
	}

	return randompost, nil
}
