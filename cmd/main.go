package main

import (
	"log"

	b "github.com/jrmanes/go-cidr/cmd/bootstrap"
)

// main entrypoint for our application
func main() {
	configFile := "./config/config.yml"

	if err := b.Run(configFile); err != nil {
		log.Fatal(err)
	}
}
