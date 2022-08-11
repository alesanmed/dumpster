// Package docker provides functions for interacting with docker engine API
package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type port struct {
	PrivatePort int64  `json:"PrivatePort"`
	PublicPort  int64  `json:"PublicPort"`
	Type        string `json:"Type"`
}

type hostConfig struct {
	NetworkMode string `json:"NetworkMode"`
}

type network struct {
	NetworkID           string `json:"NetworkID"`
	EndpointID          string `json:"EndpointID"`
	Gateway             string `json:"Gateway"`
	IPAddress           string `json:"IPAddress"`
	IPPrefixLen         int    `json:"IPPrefixLen"`
	IPv6Gateway         string `json:"IPv6Gateway"`
	GlobalIPv6Address   string `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int    `json:"GlobalIPv6PrefixLen"`
	MacAddress          string `json:"MacAddress"`
}

type networkSettings struct {
	Networks map[string]network `json:"Networks"`
}

type mountPoint struct {
	Name        string `json:"Name"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Driver      string `json:"Driver"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
	Propagation string `json:"Propagation"`
}

type filter struct {
	Ancestor  []string `json:"ancestor,omitempty"`
	Before    []string `json:"before,omitempty"`
	Expose    []string `json:"expose,omitempty"`
	Exited    []string `json:"exited,omitempty"`
	Health    []string `json:"health,omitempty"`
	Id        []string `json:"id,omitempty"`
	Isolation []string `json:"isolation,omitempty"`
	IsTask    []string `json:"is-task,omitempty"`
	Label     []string `json:"label,omitempty"`
	Name      []string `json:"name,omitempty"`
	Network   []string `json:"network,omitempty"`
	Publish   []string `json:"publish,omitempty"`
	Since     []string `json:"since,omitempty"`
	Status    []string `json:"status,omitempty"`
	Volume    []string `json:"volume,omitempty"`
}

type Container struct {
	Id              string            `json:"Id"`
	Names           []string          `json:"Names"`
	Image           string            `json:"Image"`
	ImageID         string            `json:"ImageID"`
	Command         string            `json:"Command"`
	Created         int64             `json:"Created"`
	State           string            `json:"State"`
	Status          string            `json:"Status"`
	Ports           []port            `json:"Ports"`
	Labels          map[string]string `json:"Labels"`
	SizeRw          int64             `json:"SizeRw"`
	SizeRootFs      int64             `json:"SizeRootFs"`
	HostConfig      hostConfig        `json:"HostConfig"`
	NetworkSettings networkSettings   `json:"NetworkSettings"`
	Mounts          []mountPoint      `json:"Mounts"`
}

// QueryContainers function gets all the running containers with the dumpster.enable label set to true
func QueryContainers() (*[]Container, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:2375/containers/json", nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	query, err := json.Marshal(filter{
		Label: []string{"dumpster.enable=true"},
	})

	if err != nil {
		return nil, err
	}

	q.Add("filters", string(query))

	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(res.Body)

	var containers []Container

	err = json.Unmarshal(body, &containers)

	if err != nil {
		return nil, err
	}

	fmt.Println(containers)

	return &containers, nil
}
