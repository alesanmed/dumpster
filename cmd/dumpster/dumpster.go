package main

import (
	"log"

	dumpster "github.com/alesanmed/dumpster/internal/cmd"
)

func main() {
	err := dumpster.ProcessBackups()

	if err != nil {
		log.Fatal(err)
	}
}
