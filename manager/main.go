package main

import (
	"flag"
	"fmt"
	"manager/container"
	"manager/database"
)

type Flags struct {
	Command string
	Service string
	Port		string
}

func main() {
	var flags Flags
	flags.Command = *flag.String("c", "", "Command name.")
	flags.Service = *flag.String("s", "", "Service name.")
	flags.Port = *flag.String("p", "", "Port number for connection.")
	
	flag.Parse()
	
	switch flags.Command {
	case "run":
		container.MakeAndRun(flags.Service, flags.Port)
	case "stop":
		container.StopContainer(flags.Service)
	case "gen":
		container.GenerateDockerfile(flags.Service, "FROM httpd\n\nEXPOSE 80\n\n")
		container.MakeAndRun(flags.Service, flags.Port)
	case "rm":
		container.RemoveContainerAndImage(flags.Service)
	default:
		message := "Not Found command \"%s\".\n"
		message += "You can use commands are below.\n"
		message += "-------------------\n"
		message += "run     make and run the service.\n"
		message += "stop    stop the service.\n"
		message += "gen     generate service and run.\n"
		message += "rm    	remove service.\n"

		// fmt.Printf(message, command)
	}

	kvs := database.GetAllData()
	fmt.Printf("%v", kvs)
	
}