package main

import (
	"log"

	b "github.com/jrmanes/go-cidr/cmd/bootstrap"
)

// main entrypoint for our application
func main() {
	if err := b.Run(); err != nil {
		log.Fatal(err)
	}
}
