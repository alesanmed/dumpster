package dumpster

import (
	"fmt"

	"github.com/alesanmed/dumpster/internal/docker"
)

func ProcessBackups() error {
	containers, err := docker.QueryContainers()

	fmt.Println(containers)

	// TODO: everything

	return err
}
