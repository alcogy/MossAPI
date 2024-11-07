package main

import (
	"flag"
	"log"
	"os"

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
	if os.Args[1] == "admin" {
		admin.Serve(mysql)
	} else {
		// Command
		c := flag.String("c", "", "Command name.")
		s := flag.String("s", "", "Service name.")
		a := flag.String("a", "", "Artifact directory path.")
		e := flag.String("e", "", "execute command.")

		flag.Parse()
		flags := model.Flags{Command: *c, Service: *s, Artifact: *a, Execute: *e}
		
		command.SwitchCommand(flags, mysql)
	}
	
	
}