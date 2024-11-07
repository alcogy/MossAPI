package command

import (
	"fmt"
	"io"
	"io/fs"
	"manager/container"
	"manager/model"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
)

func SwitchCommand(flags model.Flags, db *sqlx.DB) {
	switch flags.Command {
	
	case "run":
		Run(flags)

	case "stop":
		cid := container.GetContainerID(flags.Service)
		container.StopContainer(cid)

	case "gen":
		Gen(flags)

	case "rm":
		Remove(flags)
	
	case "copy":
		CopyArtifact(flags)

	default:
		message := "Not Found command \"%s\".\n"
		message += "You can use commands are below.\n"
		message += "-------------------\n"
		message += "run     make and run the service.\n"
		message += "stop    stop the service.\n"
		message += "gen     generate service and run it.\n"
		message += "rm    	remove service.\n"

		fmt.Printf(message, flags.Command)
	}
}

// ----------------------------------------
// Build image and run container.
func Run(flags model.Flags) {
	// TODO This line is for dev and debug.
	Remove(flags)

	if flags.Service == "" {
		fmt.Printf("You must specify service name.")
		return
	}
	// Check Exist Service
	container.BuildAndRun(flags.Service)
}

// ----------------------------------------
// Generate Dockerfile with content. And Run Container.
func Gen(flags model.Flags) {
	
	// Make Dockerfile with content.
	// content := container.GenerateContent(flags.Service, nil)
	// container.GenerateDockerfile(flags.Service, content)
}

// ----------------------------------------
// Remove container, docker image and service directory
func Remove(flags model.Flags) {
	container.RemoveContainerAndImage(flags.Service)

	// TODO Cooment out for dev and debug.
	// os.RemoveAll(container.GetServiceDir(flags.Service))
}

// ----------------------------------------
// Copy artifact for api to service directory.
func CopyArtifact(flags model.Flags) {
	filepath.Walk(flags.Artifact, func (path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		fmt.Println(path)

		root := "../services/" + flags.Service
		from := strings.Replace(path, flags.Artifact, "", -1)

		// Make directory
		if info.IsDir() && path != flags.Artifact {
			dist := root + from
			os.MkdirAll(dist, 0750)
		}

		if !info.IsDir() {
			r, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer r.Close()

			w, err := os.Create(root + from)
			if err != nil {
				panic(err)
			}
			defer w.Close()

			io.Copy(w, r)
		}
		return nil
	})
}
