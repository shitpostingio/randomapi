package main

import (
	"fmt"
	"log"
)

const (
	bindString string = "localhost:%d"
)

var (
	errChan chan error

	//Build is the git commit ref
	Build string
	//Version is the memesapi version
	Version string
)

func main() {

	go func() {
		for err := range errChan {
			log.Println(err)
		}
	}()

	server := Setup(fmt.Sprintf(bindString, c.Port), allowedOrigins)

	log.Printf("random memes api started\nVersion: %s\nBuild: %s", Version, Build)
	log.Fatal(server.ListenAndServe())
}
