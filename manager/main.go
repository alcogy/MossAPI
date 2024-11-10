package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"manager/admin"
	"manager/command"
	"manager/database/mysql"
)

func main() {
	// Database Open
	mysql, err := mysql.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	// Show admin on browser.
	arg := os.Args[1]
	
	if arg == "admin" {
		admin.Serve(mysql)

	} else if arg == "rm" {
		service := os.Args[2]
		fmt.Println(service)
		if service == "" {
			fmt.Println("Please specify service name.")
			return 
		}

		command.RemoveService(service, mysql)
		
	} else {
		// Command
		var f string
		flag.StringVar(&f, "f", "", "Filepath to json")
		flag.Parse()

		if f == "" {
			fmt.Println("Please specify file path with -f option.")
			return
		}
		
		command.ExecuteBuild(f, mysql)
	}
	
}