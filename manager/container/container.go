package container

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"manager/database/redis"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
)

type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Port   string `json:"port"`
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
func AllContainers() []Container {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{ All: true })
	if err != nil {
		panic(err)
	}

	services := redis.FetchAllData()
	
	var containerInfos []Container;
	for _, ctr := range containers {
		name := ctr.Names[0][1:]
		if !hasService(name, services) {
			continue
		}
		
		var port string
		if len(ctr.Ports) != 0 {
			port = strconv.FormatInt(int64(ctr.Ports[0].PublicPort), 10)
		}
		c := Container{
			ID: ctr.ID[:12],
			Name: name,
			Port: port,
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

// ----------------------------------------------------
// BuildAndRun makes docker image and create conteiner.
func BuildAndCreate(service string, port string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)
	build(ctx, cli, service)
	createContaier(ctx, cli, service, port)
}

// BuildAndRun makes docker image and run conteiner.
// Image build with tar style byte data.
func BuildAndRun(service string, port string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	build(ctx, cli, service)
	container := createContaier(ctx, cli, service, port)
	run(ctx, cli, container.ID)
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

func build(ctx context.Context, cli *client.Client, service string) {
	path := GetServiceDir(service)
	
	res, err := cli.ImageBuild(
		ctx,
		makebuildContext(path),
		types.ImageBuildOptions{
			NoCache: true,
			Remove: true,
			Tags:    []string{service},
		},
	)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	// Print from docker message.
	io.Copy(os.Stdout, res.Body)
}

func makebuildContext(root string) *bytes.Reader {
	// Make buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)

		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Make file name relative from root(args) path.
		file := strings.Replace(filepath.ToSlash(path), filepath.ToSlash(root + "\\"), "", -1)
		
		// Write header info.
		if err := tw.WriteHeader(&tar.Header{
			Name: file,
			Size: info.Size(),
			Mode: 0755,
			ModTime: info.ModTime(),
		}); err != nil {
			return err
		}

		body, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		tw.Write(body)
		// Write body data.
		// if _, err := io.Copy(tw, f); err != nil {
		// 	return err
		// }
		
		return nil
	}); err != nil {
		panic(err)
	}

	return bytes.NewReader(buf.Bytes())
}

func createContaier(ctx context.Context, cli *client.Client, service string, port string) container.CreateResponse {
	container, err := cli.ContainerCreate(
		ctx,
		&container.Config{ Image: service, ExposedPorts: nat.PortSet{"9000/tcp": struct{}{}} },
		&container.HostConfig{
			PortBindings: nat.PortMap{nat.Port("9000/tcp"): []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					HostPort: port,
				},
			}},
		},
		nil,
		nil,
		service,
	)
	if err != nil {
		panic(err)
	}

	return container
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

func hasService(name string, services []redis.KeyValue) bool {
	for _, v := range services {
		if v.Key == name {
			return true
		}
	}
	return false
}