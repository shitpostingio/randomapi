package client

import "time"

// Response represent a post.
type Response struct {
	ID    string `json:"id,omitempty"`
	Post  Data   `json:"post,omitempty"`
	Error string `json:"error,omitempty"`
}

// Data contains data info about a random post
type Data struct {
	URL      string     `json:"url"`
	Filename string     `json:"filename"`
	Type     string     `json:"mediatype"`
	Date     *time.Time `json:"date"`
}
