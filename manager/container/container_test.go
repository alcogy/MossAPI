package container

import (
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
func TestGenerateContent(t *testing.T) {
	beforeAll(t)
	
	content := GenerateContent("franc", []string{"golang-go", "dotnet-sdk-8.0"})
	GenerateDockerfile("franc", content)
	//TODO COPY binnary file to service dir.
}