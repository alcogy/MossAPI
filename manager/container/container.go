package container

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
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

func GetServiceDir(service string) string {
	return  "../services/" + service
}

// ----------------------------------------------------
// Generate Dockerfile with content to service directory.
func GenerateDockerfile(service string, content string) {
	path := GetServiceDir(service) 
	// Make service directory if not exsist.
	info, err := os.Stat(path)
	
	if err != nil || info == nil || !info.IsDir() {
		os.Mkdir(path, 0750)
	}
	// Create blank docker file.
	f, err := os.Create(path + "/Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Write content.
	_, err = f.Write([]byte(content))
	if err != nil {
		log.Fatal(err)
	}
}

// ----------------------------------------------------
// BuildAndRun makes docker image and run conteiner.
// Image build with tar style byte data.
func BuildAndRun(service string, port string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	build(&ctx, cli, service)
	run(&ctx, cli, service, port)
}

func build(ctx *context.Context, cli *client.Client, service string) {
	path, _ := filepath.Abs(GetServiceDir(service))
	
	res, err := cli.ImageBuild(
		*ctx,
		makebuildContext(path),
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{service},
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

		// Write body data.
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		
		return nil
	}); err != nil {
		panic(err)
	}

	return bytes.NewReader(buf.Bytes())
}

func run(ctx *context.Context, cli *client.Client, service string, port string) {
	res, err := cli.ContainerCreate(
		*ctx,
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

	if err := cli.ContainerStart(
		*ctx,
		res.ID,
		container.StartOptions{},
	); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(*ctx, res.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(*ctx, res.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

// ----------------------------------------------------
// Delete container and docker image.
// TODO Delete conteiner and image not perfect.
func RemoveContainerAndImage(service string) {
	StopContainer(service)
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
	for _, v := range containers {
		if v.Names[0] == "/" + service {
			cli.ContainerRemove(ctx, v.ID, container.RemoveOptions{})
			cli.ContainersPrune(ctx, filters.Args{})
		}
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
	// exec.Command("cmd", "/c", "docker", "rm", service).Output()
	// exec.Command("cmd", "/c", "docker", "rmi", service).Output()
}

// ----------------------------------------------------
// Just stop container.
func StopContainer(service string) {
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

	for _, v := range containers {
		if v.Names[0] == "/" + service {
			cli.ContainerStop(ctx, v.ID, container.StopOptions{})
			fmt.Println("Stop container.")
		}
	}
	// exec.Command("cmd", "/c", "docker", "stop", service).Output()
}

type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Port   string `json:"port"`
	Status string `json:"status"`
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

	var containerInfos []Container;
	for _, ctr := range containers {
		c := Container{
			ID: ctr.ID[:12],
			Name: ctr.Names[0],
			Port: strconv.FormatInt(int64(ctr.Ports[0].PrivatePort), 10),
			Status: ctr.Status,
		}
		containerInfos = append(containerInfos, c);
	}

	return containerInfos
}