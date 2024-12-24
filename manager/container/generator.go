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
	executes := strings.Split(body.Execute, " ")
	var cmds []string
	for _, command := range executes  {
		cmds = append(cmds, "\"" + command + "\"") 
	}
	
	content += getBaseImageAndUpdate(body.Base)
	content += "WORKDIR /app\n"
	content += "COPY . .\n\n"
	content += body.Options
	content += "\n"
	content += "EXPOSE 9000\n"
	content += "CMD ["
	content += strings.Join(cmds, ",")
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


func getBaseImageAndUpdate(base string) string {
	switch base {
	case "delian":
		return "FROM debian:12-slim\n\nRUN apt update && apt upgrade -y\n"
	
	case "centos":
		return "FROM centos:7\n\nRUN yum update && yum upgrade -y\n"
	
	case "alpine":
		return "FROM alpine:3.14\n\nRUN apk update && apk upgrade -y\n"

	case "python":
		return "FROM python\n\n"

	default:
		return "FROM debian:12-slim\n\nRUN apt update && apt upgrade -y\n"
	}
}