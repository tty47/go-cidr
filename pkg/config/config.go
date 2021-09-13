package config

import (
	"log"
	"strings"

	"github.com/jrmanes/go-cidr/pkg/ip"
	"github.com/jrmanes/go-cidr/pkg/store/memory"
)

var (
	tmp []string
	err error
)

// LoadConfiguration get an array of strings and calculate each host from the cidr
func LoadConfiguration(cidr []string) error {
	// loop inside all the cidr in order to calculate the needed hosts
	for _, n := range cidr {
		// clean array each time inside the loop
		tmp := []string{}

		// check if ip in config file is a range (contains a /X)
		if strings.Contains(n, "/") {
			// calculate the host per each ip network
			// return the ip into tmp array
			tmp, err = ip.Hosts(n)
			if err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			// is only a host - unique host
			tmp = append(tmp, n)
		}

		// store data depending on the value received
		storeData(tmp)
	}

	return nil
}

// storeData will store our data in a different paths
func storeData(ipToStore []string) {
	memory.StoreIps(ipToStore)
}
