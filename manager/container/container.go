package container

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

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
