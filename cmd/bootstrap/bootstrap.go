package bootstrap

import (
	"log"

	"github.com/jrmanes/go-cidr/pkg/config"
)

// Run function which starts everything
func Run(cidr []string) error {
	// load configuration and return an error if there is any issue
	err := config.LoadConfiguration(cidr)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return err
	}

	return err
}
