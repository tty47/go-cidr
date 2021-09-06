# Go-CIDR

[![Go Reference](https://pkg.go.dev/badge/github.com/jrmanes/go-cidr.svg)](https://pkg.go.dev/github.com/jrmanes/go-cidr)

---

## Description

The purpose of that library is check an IP inside from a CIDR.

---

## Project structure

The project structure explained what contains what.

```shell
├── cmd
│   ├── bootstrap
│   │   └── bootstrap.go // point where init the service
│   └── main.go // entrypoint of the service
├── config
│   └── config.yml // example config file
├── pkg
│   ├── config
│   │   └── config.go // config structs, methods and functions
│   ├── ip
│   │   ├── ip.go // logic to create the hosts from cirds
│   │   └── ip_test.go // test file for ip pkg
│   └── store
│       └── in_memory
│           └── in_memory.go // storage in memory
└── README.md // Readme
```

----

## Example

Using storage in memory:
```go
package main

import (
	"time"
	"log"

	b "github.com/jrmanes/go-cidr/cmd/bootstrap"
	m "github.com/jrmanes/go-cidr/pkg/store/in_memory"
)

func main() {
	// load configuration
	b.Run()

	now := time.Now()
	log.Println("---------FROM MEMORY------------")
	ip2CheckInMemory := "10.251.0.162"
	existsInMemory := m.CheckIfIpExistsInMemory(ip2CheckInMemory)
	if existsInMemory {
		log.Println("IP found ip2CheckInMemory (", ip2CheckInMemory, ")", existsInMemory)
	} else {
		log.Println("IP NOT found ip2CheckInMemory (", ip2CheckInMemory, ")", existsInMemory)
	}
	log.Println("elapse using storage memory:", time.Since(now))
	log.Println("-----------------------------")
}


```

----

Jose Ramón Mañes

---