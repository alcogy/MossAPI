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

	// TODO Implementdump.

	// Show admin on browser.
	if os.Args[1] == "admin" {
		admin.Serve(mysql)
	} else {
		// Command
		f := flag.String("f", "", "Filepath to json")
		flag.Parse()

		if *f == "" {
			fmt.Println("Please specify file path with -f option.")
		}
		
		command.ExecuteBuild(*f, mysql)
	}
	
	
}