package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shitpostingio/randomapi/rest/client"
)

// writeError writes a JSON-formatted error message on w, with err as reason,
// and with status as HTTP status.
func writeError(w http.ResponseWriter, err error, errorCode int) {
	w.WriteHeader(errorCode)

	jenc := json.NewEncoder(w)
	encErr := jenc.Encode(&client.Response{
		Error: fmt.Sprintf("could not retrieve random post: %s", err.Error()),
	})

	if encErr != nil {
		log.Printf("could not marshal error: %s\n", encErr)
	}
}
