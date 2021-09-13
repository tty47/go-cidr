package main

import (
	"log"

	b "github.com/jrmanes/go-cidr/cmd/bootstrap"
)

// main entrypoint for our application
func main() {
	// cidr to test the service
	cidr := []string{"10.251.0.0/16", "172.18.0.0/16", "192.168.0.0/16"}

	if err := b.Run(cidr); err != nil {
		log.Fatal(err)
	}
}
