package main

import (
	"flag"
	"fmt"
	"manager/container"
	"os"
)

func main() {
	command := os.Args[1]
	service := flag.String("s", "", "Service name.")
	port := flag.String("p", "12345", "Port number for connection.")

	switch command {
	case "run":
		container.MakeAndRun(*service, *port)
	case "stop":
		container.StopContainer(*service)
	case "gen":
		container.GenerateDockerfile(*service, "FROM httpd\n\nEXPOSE 80\n\n")
		container.MakeAndRun(*service, *port)
	case "rm":
		container.RemoveContainerAndImage(*service)
	default:
		message := "Not Found command \"%s\".\n"
		message += "You can use commands are below.\n"
		message += "-------------------\n"
		message += "run     make and run the service.\n"
		message += "stop    stop the service.\n"
		message += "gen     generate service and run.\n"
		message += "rm    	remove service.\n"

		fmt.Printf(message, command)
	}
	
}