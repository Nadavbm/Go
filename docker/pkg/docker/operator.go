package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
)

func (d *DockerOperator) ListContainers() error {
	if err != nil {
		d.Logger.Error(err, "could not set client")
	}
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		d.Logger.Error(err, "could not list containers")
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
	return nil
}
