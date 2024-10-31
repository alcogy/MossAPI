package main

import (
	"flag"
	"log"
	"manager/admin"
	"manager/command"
	"manager/database/mysql"
	"manager/model"
)

func main() {
	// Database Open
	mysql, err := mysql.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	// TODO for debug.
	admin.Serve(mysql)
	return

	// Command
	c := flag.String("c", "", "Command name.")
	s := flag.String("s", "", "Service name.")
	p := flag.String("p", "", "Port number for connection.")
	a := flag.String("a", "", "Artifact directory path.")
	flag.Parse()
	flags := model.Flags{Command: *c, Service: *s, Port: *p, Artifact: *a}
	
	command.SwitchCommand(flags, mysql)
}