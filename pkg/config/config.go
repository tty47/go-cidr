package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jrmanes/go-cidr/pkg/ip"
	"github.com/jrmanes/go-cidr/pkg/store/in_memory"

	"gopkg.in/yaml.v2"
)

// CIDR is the main configuration struct
type CIDR struct {
	CIDR ConntrackConfig `yaml:"cidr"`
}

// ConntrackConfig
type ConntrackConfig struct {
	Exclusions Exclusions `yaml:"exclusions"`
}

// Exclusions
type Exclusions struct {
	Src []string `yaml:"src"`
	Dst []string `yaml:"dst"`
}

func NewConfig(configFile string) (conf *CIDR, err error) {
	// check if the var contains the paths
	if configFile == "" {
		// if not, set the path to default one
		configFile = "./config/config.yml"
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	conf = &CIDR{}
	err = yaml.Unmarshal(data, conf)
	return
}

func NewCIDR() CIDR {
	return CIDR{}
}

func NewExclussions() Exclusions {
	return Exclusions{}
}

func LoadConfiguration(configFile string) error {
	// load the config file and check if there is any error
	cfg, err := NewConfig(configFile)
	//allIPS := strings.Split(d,"/")
	if err != nil {
		fmt.Println("load config fail: ", err)
		return err
	}

	// create new config
	d := cfg.CIDR
	e := d.Exclusions

	// Create a var for dst ips, and append them to the src ips' var
	srcLoop := e.Src
	dstLoop := e.Dst

	allInOneLoop := append(srcLoop, dstLoop...)

	// Uncomment to debug
	//log.Println("Var allInOneLoop: ", allInOneLoop)
	for _, s := range allInOneLoop {
		// create an empty/temporary array each time
		var tmp []string

		// check if ip in config file is a range (contains a /X)
		if strings.Contains(s, "/") {
			// calculate the host per each ip network
			// return the ip into tmp array
			tmp, err = ip.Hosts(s)
			if err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			// is only a host - unique host
			tmp = append(tmp, s)
		}

		// store data depending on the value received
		storeData(tmp)
	}

	return nil
}

// storeData will store our data in a different paths
func storeData(ipToStore []string) {
	in_memory.StoreIps(ipToStore)
}
