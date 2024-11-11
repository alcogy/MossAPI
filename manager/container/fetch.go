package container

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
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
	containers, err := fetchContainers()
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

func FetchAllServicesFull() ([]ContainerFull, error) {
	containers, err := fetchContainers()
	if err != nil {
		return nil, err
	}
	
	var containerInfos []ContainerFull;
	for _, ctr := range containers { 
		if ctr.Labels["group"] != "service" {
			continue
		}

		c := ContainerFull{
			ID: ctr.ID,
			Names: ctr.Names,
			ImageID: ctr.ImageID,
			Command: ctr.Command,
			Created: ctr.Created,
			State: ctr.State,
			Status: ctr.Status,
		}
		containerInfos = append(containerInfos, c);
	}

	return containerInfos, nil
}

func fetchContainers() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{ All: true })
	if err != nil {
		return nil, err
	}

	return containers, nil
}