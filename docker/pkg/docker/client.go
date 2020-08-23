package main

import (
	"docker.io/go-docker"
	"github.com/docker/docker/client"
	"go.uber.org/zap"
)

type DockerOperator struct {
	Logger *zap.Logger
	Client *docker.Client
}

func NewClient(logger *zap.Logger) (*DockerOperator, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &DockerOperator{
		Logger:	logger,
		Client: *cli
	}
}
