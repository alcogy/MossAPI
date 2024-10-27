package container

import (
	"fmt"
	"testing"
)

func TestGenerateContent(t *testing.T) {
	content := GenerateContent("franc", "9000", nil)
	fmt.Println(content)
	// GenerateDockerfile("franc", content)
	//TODO COPY binnary file to service dir.
}