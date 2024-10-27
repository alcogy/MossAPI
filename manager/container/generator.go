// Generate Dockerfile Content.
package container

import "strings"

func GenerateContent(service string, addPackage []string) string {
	var content string
	pkgs := strings.Join(addPackage, " ")

	// TODO make template file.
	content += "FROM ubuntu:latest\n\n"
	content += "RUN apt update && apt upgrade -y\n"
	if pkgs != "" {
		content += "RUN apt install -y " + pkgs + "\n\n"
	}
	content += "WORKDIR /app\n"
	content += "EXPOSE 9000\n"
	content += "COPY . .\n\n"
	content += "CMD [\"./" + service + "\"]"

	return content
}