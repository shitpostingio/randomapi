package main

import (
	"fmt"
	"log"

	"github.com/shitposting/randomapi/backstore"
	"github.com/shitposting/randomapi/rest"
)

const (
	bindString string = "localhost:%d"
)

var (
	feedErrors chan error
	bs         *backstore.Backstore

	//Build is the git commit ref
	Build string
	//Version is the memesapi version
	Version string
)

func main() {

	go func() {
		for err := range feedErrors {
			log.Println(err)
		}
	}()

	server := rest.Setup(fmt.Sprintf(bindString, c.Port), bs, allowedOrigins)

	log.Printf("random memes api started\nVersion: %s\nBuild: %s", Version, Build)
	log.Fatal(server.ListenAndServe())
}
