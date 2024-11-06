package container

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// GetContainerID get container ID by service name.
func GetContainerID(service string) string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var containerID string
	for _, v := range containers {
		if v.Names[0] == "/" + service {
			containerID = v.ID
		}
	}

	return containerID
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
			Status: ctr.Status,
		}
		containerInfos = append(containerInfos, c);
	}

	return containerInfos
}

// ----------------------------------------------------
// Run starts docker container and create conteiner.
func Run(conteinerID string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	run(ctx, cli, conteinerID)
}

// Just stop container.
func StopContainer(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctx := context.Background()
	
	cli.ContainerStop(ctx, containerID, container.StopOptions{})
	fmt.Println("Stoped container.")
}

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
// TODO Delete conteiner and image not perfect.
func RemoveContainerAndImage(service string) {
	containerID := GetContainerID(service)
	
	StopContainer(containerID)
	
	// Get Client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctx := context.Background()
	
	if (containerID != "") {
		cli.ContainerRemove(ctx, containerID, container.RemoveOptions{})
		cli.ContainersPrune(ctx, filters.Args{})
	}
	
	images, err := cli.ImageList(ctx, image.ListOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range images {
		if len(img.RepoTags) == 0 {
			continue
		}
		tags := strings.Split(img.RepoTags[0], ":")
		if tags[0] == service {
			fmt.Println(img.ID)
			cli.ImageRemove(ctx, img.ID, image.RemoveOptions{Force: true, PruneChildren: true})
			cli.ImagesPrune(ctx, filters.Args{})
		}
	}
	
	fmt.Println("Remove container and image.")
}

func run(ctx context.Context, cli *client.Client, containerID string) {
	
	if err := cli.ContainerStart(
		ctx,
		containerID,
		container.StartOptions{},
	); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
