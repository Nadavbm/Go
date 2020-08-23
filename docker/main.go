package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	err := listContainers()
	if err != nil {
		log.Println("ERROR:", err)
	}
}

func listContainers() error {
	config, err := setConfig()
	if err != nil {
		return err
	}
	containers, err := config.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}

	fmt.Println("docker containers:\nid\t\tname\t\tstatus")
	for _, container := range containers {
		fmt.Printf("%s\t%s\t%s\n", container.ID[:10], container.Image, container.Status)
	}

	return nil
}

func setConfig() (*client.Client, error) {
	config, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return config, nil
}
