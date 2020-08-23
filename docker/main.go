package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	containers, err := ListContainers()
	if err != nil {
		log.Println("ERROR: listing docker container:", err)
	}
	for _, container := range containers {
		fmt.Println(string(container.ID[:12]))
	}

	err = CreateContainer()
	if err != nil {
		log.Println("ERROR: creating container:", err)
	}

	containers, err = ListContainers()
	if err != nil {
		log.Println("ERROR: listing docker container:", err)
	}
	for _, container := range containers {
		fmt.Println(string(container.ID[:10]))
	}

	err = GetContainerLogs()
	if err != nil {
		log.Println("ERROR: cannot get container logs")
	}

	err = StopAllRunningContainers()
	if err != nil {
		log.Println("ERROR: creating container:", err)
	}

	containers, err = ListContainers()
	if err != nil {
		log.Println("ERROR: listing docker container:", err)
	}
	for _, container := range containers {
		fmt.Println(string(container.ID[:10]))
	}
}

func SetContainerConfig() error {
	//config, err := client.ContainerAPIClient()
	return nil
}

func SetConfig() (*client.Client, error) {
	config, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return config, nil
}

// list all docker containers - docker ps
func ListContainers() ([]types.Container, error) {
	config, err := SetConfig()
	if err != nil {
		return nil, err
	}

	containers, err := config.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	fmt.Println("\nRUNNING CONTAINERS:\nID\t\tNAME\t\tSTATUS")
	for _, container := range containers {
		fmt.Printf("%s\t%s\t%s\n", container.ID[:10], container.Image, container.Status)
	}

	return containers, nil
}

// create a docker container
func CreateContainer() error {
	cli, err := SetConfig()
	if err != nil {
		log.Println(err, "")
	}

	reader, err := cli.ImagePull(context.Background(), "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		log.Println(err, "")
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   false,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Println(err, "start container error")
	}

	fmt.Println(resp.ID)
	return nil
}

func GetContainerLogs() error {
	ctx := context.Background()
	cli, err := SetConfig()
	if err != nil {
		log.Println("could not set config for docker")
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	containers, err := ListContainers()
	if err != nil {
		log.Println("could not list containers:", err)
	}

	fmt.Println("\nALL CONTAINER LOGS:")
	for _, container := range containers {
		out, err := cli.ContainerLogs(ctx, string(container.ID[:12]), options)
		if err != nil {
			log.Println("", err)
		}

		io.Copy(os.Stdout, out)
	}
	return nil
}

func StopAllRunningContainers() error {
	ctx := context.Background()
	cli, err := SetConfig()
	if err != nil {
		log.Println(err, "set config error")
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Println(err, "listing container error")
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			log.Println(err, "stopping container error")
		}
		fmt.Println("Success")
	}
	return nil
}
