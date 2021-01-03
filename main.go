package main

import (
	"fmt"
	"log"
)

const (
	bindString string = "localhost:%d"
)

var (
	//Build is the git commit ref
	Build string
	//Version is the randomapi version
	Version string
)

func main() {
	go startCleanRequestRoutine()

	server := Setup(fmt.Sprintf(bindString, c.Port), allowedOrigins)

	log.Printf("random api started\nVersion: %s Build: %s", Version, Build)
	log.Fatal(server.ListenAndServe())
}
