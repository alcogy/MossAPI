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
	State string `json:"state"`
	Status string `json:"status"`
}

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

func run(ctx context.Context, cli *client.Client, containerID string) error {
	
	if err := cli.ContainerStart(
		ctx,
		containerID,
		container.StartOptions{},
	); err != nil {
		return err
	}

	go func() {
		statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				return 
			}
		case <-statusCh:
		}
	
		out, err := cli.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true})
		if err != nil {
			panic(err)
		}
	
		stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	}()
	
	return nil
}
