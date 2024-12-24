package command

import (
	"fmt"
	"manager/table"
	"os"
	"testing"
)

func beforeAll(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(cwd)
	})
	os.Chdir("../")
}

func TestDump(t *testing.T) {
	beforeAll(t)
	db, _ := table.Connection()
	
	err := Dump("./test.json", db)
	if err != nil {
		t.Fatal(err.Error())
	}
	os.RemoveAll("./test.json")
}

func TestFetchAllTables(t *testing.T) {
	beforeAll(t)
	db, _ := table.Connection()
	tables := fetchAllTables(db)

	if len(tables) == 0 {
		t.Fatal("Couldn't get tables")
	}

	if len(tables[0].Columns) == 0 {
		t.Fatal("Couldn't get columns")
	}

	fmt.Println(tables)
}