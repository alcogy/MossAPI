package container

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func beforeAll(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(cwd)
	})
	os.Chdir("../")
}
func TestGetServiceDir(t *testing.T) {
	beforeAll(t)
	
	path := GetServiceDir("customer")
	if path == "" {
		t.Fatal("path isn't get.")
	}

	fmt.Println(path)
}

func TestGetContainerID(t *testing.T) {
	beforeAll(t)

	cid := GetContainerID("customer")
	want := "763ad7604424"
	if cid[:12] != want {
		t.Fatalf("Expected %v, but get %v", want, cid[:12])
	}
	
	fmt.Println(cid)
}

func TestGenerateDockerfile(t *testing.T) {
	beforeAll(t)

	service := "gotest"
	content := "FROM apache:latest\nRUN echo 'Hello'"
	GenerateDockerfile(service, content)
	
	path := GetServiceDir(filepath.Join(service, "Dockerfile")) 
	_, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	os.RemoveAll(GetServiceDir(service))
}

func TestAllContainers(t *testing.T) {
	beforeAll(t)

	containers := AllContainers()
	if len(containers) == 0 {
		t.Fatal("Container not found.")
	}

	fmt.Println(containers)
}

func TestRemove(t *testing.T) {
	beforeAll(t)

	cid := "763ad7604424"
	Remove(cid)

	containers := AllContainers()
	for _, v := range containers {
		if v.ID[:12] == cid {
			t.Fatal("Container is not deleted.")
		}
	}
}

func TestBuildAndCreate(t *testing.T) {
	beforeAll(t)

	service := "buildtest"
	content := "FROM httpd\nRUN echo 'Hello'"
	GenerateDockerfile(service, content)
	BuildAndCreate(service, "11111")

	os.RemoveAll(GetServiceDir(service))
}