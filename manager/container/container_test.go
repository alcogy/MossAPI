package container

import (
	"fmt"
	"os"
	"testing"
)

func beforeAll(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(cwd)
	})
	os.Chdir("../")
}

func TestGetContainerID(t *testing.T) {
	beforeAll(t)
	service := "testservice"	
	BuildAndCreate(service)
	defer RemoveContainerAndImage(service)

	cid := GetContainerID(service)
	if cid == "" {
		t.Fatalf("ContainerID is blank")
	}
	
	fmt.Println(cid)
}

func TestAllContainers(t *testing.T) {
	beforeAll(t)
	service := "testservice"	
	BuildAndCreate(service)
	defer RemoveContainerAndImage(service)

	containers := FetchAllServices()
	if len(containers) == 0 {
		t.Fatal("Container not found.")
	}

	fmt.Println(containers)
}

func TestRemove(t *testing.T) {
	beforeAll(t)
	service := "testservice"	
	BuildAndCreate(service)
	defer RemoveContainerAndImage(service)
	
	containers := FetchAllServices()
	for _, v := range containers {
		if v.Name == "/" + service {
			t.Fatal("Container is not deleted.")
		}
	}

}

func TestBuildAndCreate(t *testing.T) {
	beforeAll(t)

	service := "buildtest"
	content := "FROM httpd\nRUN echo 'so'\nRUN echo 'some'\nRUN echo 'Hello'"
	GenerateDockerfile(service, content)
	BuildAndCreate(service)

	os.RemoveAll(GetServiceDir(service))
}
