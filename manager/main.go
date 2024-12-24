package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"manager/admin"
	"manager/command"
	"manager/table"
)

func main() {
	// Database Open
	db, err := table.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Show admin on browser.
	arg := os.Args[1]
	
	if arg == "admin" {
		admin.Serve(db)

	} else if arg == "rm" {
		service := os.Args[2]
		if service == "" {
			fmt.Println("Please specify service name.")
			return 
		}

		command.RemoveService(service, db)
	
	} else if arg == "dump" {
		path := os.Args[2]
		if path == "" {
			fmt.Println("Please export file path.")
			return
		}

		command.Dump(path, db)
	
	} else {
		// Command
		var f string
		flag.StringVar(&f, "f", "", "Filepath to json")
		flag.Parse()

		if f == "" {
			fmt.Println("Please specify file path with -f option.")
			return
		}
		
		command.ExecuteBuild(f, db)
	}
	
}