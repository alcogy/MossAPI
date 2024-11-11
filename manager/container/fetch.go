package container

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// GetContainerID get container ID by service name.
func GetContainerID(service string) string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	defer cli.Close()
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	for _, v := range containers {
		if v.Names[0] == "/" + service {
			return v.ID
		}
	}

	return ""
}

// Confirm active gateway container. 
func IsActiveGateway() bool {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		return false
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{ All: true })
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, ctr := range containers {
		name := ctr.Names[0][1:] 
		if name == "gateway" {
			if ctr.State == "running" {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// ----------------------------------------------------
// Fetch all containers.
func FetchAllServices() []Container {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{ All: true })
	if err != nil {
		panic(err)
	}
	
	var containerInfos []Container;
	for _, ctr := range containers {
		name := ctr.Names[0][1:] 
		if ctr.Labels["group"] != "service" {
			continue
		}
		c := Container{
			ID: ctr.ID[:12],
			Name: name,
			State: ctr.State,
			Status: ctr.Status,
		}
		containerInfos = append(containerInfos, c);
	}

	return containerInfos
}
