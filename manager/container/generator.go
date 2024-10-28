// Generate Dockerfile Content.
package container

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