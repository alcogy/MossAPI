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
	if err != nil || !info.IsDir() {
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
func MakeAndRun(service string, port string) {
	fmt.Println("Make container " + service)
	make(service)
	fmt.Println("Run container " + service)
	run(service, port)
}

func make(service string) {	
	path := "../services/" + service
	exec.Command("cmd", "/c", "docker", "build", "-t", service, path).Output()
}

func run(service string, port string) {
	connect := port + ":80"
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