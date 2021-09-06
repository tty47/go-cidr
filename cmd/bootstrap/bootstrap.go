package bootstrap

import (
	"log"

	"github.com/jrmanes/go-cidr/pkg/config"
)

// Run function which starts everything
func Run() error {
	// load configuration and return an error if there is any issue
	err := config.LoadConfiguration()
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return err
	}

	return err
}
