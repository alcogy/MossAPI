package main

import (
	"flag"
	"manager/command"
	"manager/model"
)

func main() {
	c := flag.String("c", "", "Command name.")
	s := flag.String("s", "", "Service name.")
	p := flag.String("p", "", "Port number for connection.")
	a := flag.String("a", "", "Artifact directory path.")
	flag.Parse()

	flags := model.Flags{Command: *c, Service: *s, Port: *p, Artifact: *a}
	
	command.SwitchCommand(flags)
	
}