package container

import (
	"fmt"
	"manager/admin/types"
	"os"
	"path/filepath"
	"testing"
)

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

func TestGenerateContent(t *testing.T) {
	beforeAll(t)
	body := types.CreateServiceBody{
		Service: "testservice",
		Options: "",
		Artifact: "/app/output",
		Execute: "python3 ./main.py",
	}

	cont := GenerateContent(body)
	if cont == "" {
		t.Fatal("Couldn't work")
	}
	fmt.Println(cont)
}