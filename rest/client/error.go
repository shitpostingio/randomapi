package client

import "encoding/json"

// Error represents an error occurred during request handling.
type Error struct {
	Error string `json:"error,omitempty"`
}

// ErrorJSON creates and marshals an Error with the given message.
func ErrorJSON(message string) []byte {
	var e Error
	e.Error = message
	j, _ := json.Marshal(e)

	return j
}
