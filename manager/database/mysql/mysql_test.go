package mysql

import (
	"fmt"
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
		TableName: "test1_table",
		TableDesc: "comment test Hello",
		Columns: []Column{
			{
				Name:  "id",
				Type:  "int",
				PK:    true,
				NotNull: true,
				Index: 0,
				Unique: 0,
				Comment: "",
			},
			{
				Name:  "name",
				Type:  "varchar(255)",
				PK:    false,
				NotNull: true,
				Index: 1,
				Unique: 1,
				Comment: "mysqltete",
			},
			{
				Name:  "area_id",
				Type:  "int",
				PK:    false,
				NotNull: false,
				Index: 3,
				Unique: 1,
				Comment: "flow tere",
			},
		},
	}

	sql := makeCreateTableSql(table)
	fmt.Println(sql)

	err = CreateTable(db, table)
	if err != nil {
		t.Fatal(err)
	}	
	
	// DeleteTable(db, table.TableName)
}