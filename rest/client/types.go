package client

import "time"

// Response represent a meme.
type Response struct {
	ID    string `json:"id,omitempty"`
	Meme  Data   `json:"meme,omitempty"`
	Error string `json:"error,omitempty"`
}

// Data contains data info about a random meme
type Data struct {
	URL      string     `json:"url"`
	Filename string     `json:"filename"`
	Type     string     `json:"mediatype"`
	Date     *time.Time `json:"date"`
}
