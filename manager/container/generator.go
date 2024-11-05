// Generate Dockerfile Content.
package container

import (
	"log"
	"os"
	"path/filepath"
)

// Generate content for Dockerfile.
// Currently base on debian:12-slim.
func GenerateContent(service string, optionalRun []string) string {
	var content string

	// TODO make template file.
	content += "FROM debian:12-slim\n\n"
	content += "RUN apt update && apt upgrade -y\n"
	for _, run := range optionalRun {
		content += "RUN " + run + "\n"
	}
	content += "\n"
	content += "WORKDIR /app\n"
	content += "EXPOSE 9000\n"
	content += "COPY . .\n\n"
	content += "CMD [\"./" + service + "\"]"

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