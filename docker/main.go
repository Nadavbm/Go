package main

import (
	dock "github.com/nadavbm/go/docker/pkg/docker"
)

func main() {
	var operator dock.DockerOperator

	operator.ListContainers()
}
