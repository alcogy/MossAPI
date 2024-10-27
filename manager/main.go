package main

import (
	"flag"
	"fmt"
	"manager/container"
	"manager/database/redis"
)

type Flags struct {
	Command string
	Service string
	Port		string
}

func main() {

	c := flag.String("c", "", "Command name.")
	s := flag.String("s", "", "Service name.")
	p := flag.String("p", "", "Port number for connection.")	
	flag.Parse()

	flags := Flags{*c, *s, *p}
	
	switch flags.Command {
	case "check":
		port, err := redis.GetPort(flags.Service)
		if err != nil || port == "" {
			fmt.Println("not exists")
		} else {
			fmt.Println("exists")
		}
		
	case "run":
		// Check Exist Service
		port, err := redis.GetPort(flags.Service)
		if err != nil && port != "" {
			fmt.Printf("%v is not exist", flags.Service)
		} else {
			container.BuildAndRun(flags.Service, port)
		}

	case "stop":
		container.StopContainer(flags.Service)

	case "gen":
		// Check port number from command args.
		if flags.Service == "" || flags.Port == "" {
			fmt.Printf("You must specify service name and port number.")
			return
		}

		if !redis.CheckPortNumberFree(flags.Port) {
			fmt.Printf("%v port is not free.", flags.Port)
			return
		}

		//Check exist service
		if redis.CheckExistService(flags.Service) {
			fmt.Printf("%v is exist", flags.Service)
			return
		}
		
		container.GenerateDockerfile(flags.Service, "FROM httpd\n\nEXPOSE 80\n\n")
		container.BuildAndRun(flags.Service, flags.Port)

	case "rm":
		container.RemoveContainerAndImage(flags.Service)
		redis.DeleteService(flags.Service)

	default:
		message := "Not Found command \"%s\".\n"
		message += "You can use commands are below.\n"
		message += "-------------------\n"
		message += "run     make and run the service.\n"
		message += "stop    stop the service.\n"
		message += "gen     generate service and run.\n"
		message += "rm    	remove service.\n"

		fmt.Printf(message, flags.Command)
	}
	
}