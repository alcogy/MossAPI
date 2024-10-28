package command

import (
	"fmt"
	"io"
	"io/fs"
	"manager/container"
	"manager/database/mysql"
	"manager/database/redis"
	"manager/model"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
)

func SwitchCommand(flags model.Flags, db *sqlx.DB) {
	switch flags.Command {
	case "check":
		Check(flags)

	case "run":
		Run(flags)

	case "stop":
		container.StopContainer(flags.Service)

	case "gen":
		Gen(flags)

	case "rm":
		Remove(flags)
	
	case "copy":
		CopyArtifact(flags)

	case "db":
		DoDB(flags, db)
	
	default:
		message := "Not Found command \"%s\".\n"
		message += "You can use commands are below.\n"
		message += "-------------------\n"
		message += "check   Check exist the service.\n"
		message += "run     make and run the service.\n"
		message += "stop    stop the service.\n"
		message += "gen     generate service and run it.\n"
		message += "rm    	remove service.\n"

		fmt.Printf(message, flags.Command)
	}
}

// ----------------------------------------
// Check exist service
func Check(flags model.Flags) {
	port, err := redis.GetPort(flags.Service)
	if err != nil || port == "" {
		fmt.Println("not exists")
	} else {
		fmt.Println("exists")
	}
}

// ----------------------------------------
// Build image and run container.
func Run(flags model.Flags) {
	if flags.Service == "" || flags.Port == "" {
		fmt.Printf("You must specify service name and port number.")
		return
	}
	// Check Exist Service
	port, err := redis.GetPort(flags.Service)
	if err != nil && port != "" {
		fmt.Printf("%v is not exist", flags.Service)
	} else {
		container.BuildAndRun(flags.Service, port)
	}
}

// ----------------------------------------
// Generate Dockerfile with content. And Run Container.
func Gen(flags model.Flags) {
	// Check port number from command args.
	if flags.Service == "" {
		fmt.Printf("You must specify service name.")
		return
	}

	if !redis.CheckPortNumberFree(flags.Port) {
		fmt.Printf("%v port is not free.", flags.Port)
		return
	}

	// Check exist service
	if redis.CheckExistService(flags.Service) {
		fmt.Printf("%v is exist", flags.Service)
		return
	}

	// Make Dockerfile with content.
	content := container.GenerateContent(flags.Service, nil)
	container.GenerateDockerfile(flags.Service, content)
}

// ----------------------------------------
// Remove container, docker image and service directory
func Remove(flags model.Flags) {
	container.RemoveContainerAndImage(flags.Service)
	redis.DeleteService(flags.Service)
	os.RemoveAll(container.GetServiceDir(flags.Service))
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

// ----------------------------------------
// Experimental DB actions. TODO: Delete
func DoDB(flags model.Flags, db *sqlx.DB) {
	// TODO Implement generate DDL.
	//mysql.ExecuteDDL(db, "CREATE TABLE `gett` (`id` int not null, `name` varchar(255) not null);")
	tables := mysql.FetchAllTable(db)
	fmt.Printf("%v", tables)
}