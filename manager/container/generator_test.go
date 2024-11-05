package container

import (
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