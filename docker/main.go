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
	err := ListContainers()
	if err != nil {
		log.Println("ERROR: listing docker container:", err)
	}

	err = ListVolumes()
	if err != nil {
		log.Println("ERROR: listing docker voluems:", err)
	}

	err = CreateContainer()
	if err != nil {
		log.Println("ERROR: creating container:", err)
	}

	err = StopAllRunningContainers()
	if err != nil {
		log.Println("ERROR: creating container:", err)
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
func ListContainers() error {
	config, err := SetConfig()
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