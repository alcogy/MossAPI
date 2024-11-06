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

	cid := GetContainerID("customer")
	want := "763ad7604424"
	if cid[:12] != want {
		t.Fatalf("Expected %v, but get %v", want, cid[:12])
	}
	
	fmt.Println(cid)
}

func TestAllContainers(t *testing.T) {
	beforeAll(t)

	containers := FetchAllServices()
	if len(containers) == 0 {
		t.Fatal("Container not found.")
	}

	fmt.Println(containers)
}

func TestRemove(t *testing.T) {
	beforeAll(t)

	cid := "763ad7604424"
	Remove(cid)

	containers := FetchAllServices()
	for _, v := range containers {
		if v.ID[:12] == cid {
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

	// ctx := context.Background()

	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// build(ctx, cli, service)
	os.RemoveAll(GetServiceDir(service))
}
