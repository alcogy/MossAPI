// Generate Dockerfile Content.
package container

import (
	"log"
	"manager/admin/types"
	"os"
	"path/filepath"
	"strings"
)

// Generate content for Dockerfile.
// Currently base on debian:12-slim.
func GenerateContent(body types.CreateServiceBody) string {
	var content string
	executes :=strings.Split(body.Execute, " ")

	// TODO make template file.
	content += "FROM debian:12-slim\n\n"
	content += "RUN apt update && apt upgrade -y\n"
	content += body.Options
	content += "\n"
	content += "WORKDIR /app\n"
	content += "EXPOSE 9000\n"
	content += "COPY . .\n\n"
	content += "CMD ["
	for _, command := range executes  {
		content += "\"" + command + "\""
	}
	content += "]"

	return content
}


// Generate Dockerfile with content to service directory.
func GenerateDockerfile(service string, content string) {
	path := GetServiceDir(service) 

	// Make service directory if not exsist.
	info, err := os.Stat(path)
	
	if err != nil || info == nil || !info.IsDir() {
		os.Mkdir(path, 0750)
	}
	// Create blank docker file.
	f, err := os.Create(filepath.Join(path, "Dockerfile"))
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