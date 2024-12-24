package container

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"manager/model"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// ----------------------------------------------------
// BuildAndRun makes docker image and create conteiner.
func BuildAndCreate(service string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)
	build(ctx, cli, service)
	createContaier(ctx, cli, service)
}

// BuildAndRun makes docker image and run conteiner.
// Image build with tar style byte data.
func BuildAndRun(service string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	build(ctx, cli, service)
	container := createContaier(ctx, cli, service)
	run(ctx, cli, container.ID)
}

func build(ctx context.Context, cli *client.Client, service string) {
	path := GetServiceDir(service)

	res, err := cli.ImageBuild(
		ctx,
		makebuildContext(path),
		types.ImageBuildOptions{
			NoCache: true,
			Remove:  true,
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
		file := strings.Replace(filepath.ToSlash(path), filepath.ToSlash(root+"\\"), "", -1)

		// Write header info.
		if err := tw.WriteHeader(&tar.Header{
			Name:    file,
			Size:    info.Size(),
			Mode:    0755,
			ModTime: info.ModTime(),
		}); err != nil {
			return err
		}

		// Write(copy) body data.
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}

	return bytes.NewReader(buf.Bytes())
}

func createContaier(ctx context.Context, cli *client.Client, service string) container.CreateResponse {
	conf := &container.Config{
		Image: service,
		ExposedPorts: nat.PortSet{"9000/tcp": struct{}{}},
		Labels: map[string]string{"group": "service"},
	}

	container, err := cli.ContainerCreate(
		ctx,
		conf,
		&container.HostConfig{},
		nil,
		nil,
		service,
	)
	if err != nil {
		panic(err)
	}

	cli.NetworkConnect(ctx, model.PrivateNetworkName, container.ID, &network.EndpointSettings{})

	return container
}