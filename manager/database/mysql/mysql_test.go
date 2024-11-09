package mysql

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func beforeAll(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(cwd)
	})
	os.Chdir("../../")
}

func sampleData(name string) Table {
	table := Table{
		TableName: name,
		TableDesc: "comment test Hello",
		Columns: []Column{
			{
				Name:  "id",
				Type:  "int",
				PK:    true,
				Nullable: false,
				Index: 0,
				Unique: 0,
				Comment: "",
			},
			{
				Name:  "name",
				Type:  "varchar(255)",
				PK:    false,
				Nullable: false,
				Index: 1,
				Unique: 1,
				Comment: "mysqltete",
			},
			{
				Name:  "area_id",
				Type:  "int",
				PK:    false,
				Nullable: false,
				Index: 3,
				Unique: 1,
				Comment: "flow tere",
			},
		},
	}
	return table
}

func create(db *sqlx.DB, tb string) {
	table := sampleData(tb)
	CreateTable(db, table)
}

func TestCreateTable(t *testing.T) {
	beforeAll(t)
	tableName := "test1_table"

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	table := sampleData(tableName)

	sql := makeCreateTableSql(table)
	fmt.Println(sql)

	err = CreateTable(db, table)
	if err != nil {
		t.Fatal(err)
	}	
	
	//DeleteTable(db, tableName)
}
