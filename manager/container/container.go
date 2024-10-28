package container

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
// Make docker image from Dockerfile and run conteiner.
func BuildAndRun(service string, port string) {
	build(service)
	run(service, port)
}

func build(service string) {
	
	fmt.Println("Make container " + service)
	path := GetServiceDir(service)
	fmt.Println(path)
	out, err := exec.Command("cmd", "/c", "docker", "build", "-t", service, path).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func run(service string, port string) {
	fmt.Println("Run container " + service)
	connect := port + ":9000"
	exec.Command("cmd", "/c", "docker", "run", "-p", connect, "--name", service, "-d", service).Output()
}

// ----------------------------------------------------
// Delete container and docker image.
func RemoveContainerAndImage(service string) {
	StopContainer(service)
	exec.Command("cmd", "/c", "docker", "rm", service).Output()
	exec.Command("cmd", "/c", "docker", "rmi", service).Output()
}

// ----------------------------------------------------
// Just stop container.
func StopContainer(service string) {
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	panic(err)
	// }
	// defer cli.Close()
	// ctx := context.Background()
	// cli.ContainerStop(ctx, service, container.StopOptions{})
	exec.Command("cmd", "/c", "docker", "stop", service).Output()
}
