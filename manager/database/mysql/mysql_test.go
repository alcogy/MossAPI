package mysql

import (
	"log"
	"os"
	"testing"
)

func beforeAll(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(cwd)
	})
	os.Chdir("../../")
}

func TestCreateTable(t *testing.T) {
	beforeAll(t)

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	table := Table{
		Name: "sample_table",
		Columns: []Column{
			{
				Name:  "id",
				Type:  "int",
				PK:    true,
				Index: false,
			},
			{
				Name:  "name",
				Type:  "varchar(255)",
				PK:    true,
				Index: false,
			},
			{
				Name:  "area_id",
				Type:  "int",
				PK:    false,
				Index: true,
			},
		},
	}

	err = CreateTable(db, table)
	if err != nil {
		t.Fatal(err)
	}	

	DeleteTable(db, table.Name)
}