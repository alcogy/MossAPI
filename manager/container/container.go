package container

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// ----------------------------------------------------
// Generate Dockerfile with content to service directory.
func GenerateDockerfile(service string, content string) {
	path := "../services/" + service
	
	// Make service directory if not exsist.
	info, err := os.Stat(path)
	fmt.Println(info)
	if err != nil || info == nil || !info.IsDir() {
		os.Mkdir(path, 0750)
	}
	fmt.Println(path)
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
	path := "../services/" + service
	exec.Command("cmd", "/c", "docker", "build", "-t", service, path).Output()
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
	exec.Command("cmd", "/c", "docker", "stop", service).Output()
}