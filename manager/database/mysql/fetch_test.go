package mysql

import (
	"fmt"
	"log"
	"testing"
)

func TestGetTableInfo(t *testing.T) {
	beforeAll(t)
	tableName := "test1_table"
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	create(db, tableName)
	defer func() {
		DeleteTable(db, tableName)
		db.Close()
	}()
	sample := sampleData(tableName)

	tableInfo := getTableInfo(db, tableName)
	fmt.Println(tableInfo)
	
	if tableInfo.Name != sample.TableName {
		t.Fatalf("Expected '%s' but got '%s'", tableName, tableInfo.Name)
	}

	if tableInfo.Comment != sample.TableDesc {
		t.Fatalf("Expected '%s' but got '%s'", sample.TableDesc, tableInfo.Comment)
	}
	
}

func TestGetColumnInfo(t *testing.T) {
	beforeAll(t)
	tableName := "test1_table"
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	create(db, tableName)
	// sample := sampleData(tableName)
	columns := getColumnInfo(db, tableName)
	fmt.Println(columns)
	if len(columns) == 0 {
		t.Fatal("couldn't get data")
	}

	DeleteTable(db, tableName)
}

func TestGetIndexInfo(t *testing.T) {
	beforeAll(t)
	tableName := "test1_table"
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	create(db, tableName)
	indexes := getIndexInfo(db, tableName)
	fmt.Println(indexes)
	for _, v := range indexes {
		if v.Table != tableName {
			t.Fatal("Couldn't get data correctly.")
		}
	}

	DeleteTable(db, tableName)
}