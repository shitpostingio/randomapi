package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func serveMeme(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	filename, ok := vars["id"]
	if !ok {
		writeError(writer, fmt.Errorf("not found"), http.StatusNotFound)
		return
	}

	post, ok := requestedPosts[filename]
	if !ok {
		writeError(writer, fmt.Errorf("not found"), http.StatusNotFound)
		return
	}

	file, err := os.Open(post.path)
	if err != nil {
		writeError(writer, err, http.StatusInternalServerError)
		return
	}

	defer func() {
		err = file.Close()
		writeError(writer, err, http.StatusInternalServerError)
		return
	}()

	if post.mediatype == "image" { // <-- set the content-type header
		writer.Header().Set("Content-Type", "image/jpeg")
	} else {
		writer.Header().Set("Content-Type", "video/mp4")
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		writeError(writer, err, http.StatusInternalServerError)
		return
	}

	delete(requestedPosts, filename)

	return
}
