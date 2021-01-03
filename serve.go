package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func servePost(writer http.ResponseWriter, request *http.Request) {

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
		if err != nil {
			log.Printf("unable to close file %v\n", err)
		}
		return
	}()

	_, err = io.Copy(writer, file)
	if err != nil {
		writeError(writer, err, http.StatusInternalServerError)
		return
	}

	delete(requestedPosts, filename)

	return
}
