package container

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// Remove is delete only container
func Remove(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctx := context.Background()

	StopContainer(containerID)
	cli.ContainerRemove(ctx, containerID, container.RemoveOptions{})
	cli.ContainersPrune(ctx, filters.Args{})
}

// ----------------------------------------------------
// Delete container and docker image.
func RemoveContainerAndImage(service string) error {
	containerID := GetContainerID(service)
	
	StopContainer(containerID)
	
	// Get Client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer cli.Close()
	ctx := context.Background()
	
	if (containerID != "") {
		cli.ContainerRemove(ctx, containerID, container.RemoveOptions{})
	}
	
	images, err := cli.ImageList(ctx, image.ListOptions{All: false})
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, img := range images {
		if len(img.RepoTags) == 0 {
			continue
		}
		tags := strings.Split(img.RepoTags[0], ":")
		if tags[0] == service {
			fmt.Println("Remove image: " + img.ID)
			cli.ImageRemove(ctx, img.ID, image.RemoveOptions{Force: true, PruneChildren: true})
			cli.ImagesPrune(ctx, filters.Args{})
		}
	}
	
	fmt.Println("Remove container and image.")

	return nil
}