package command

import (
	"encoding/json"
	"fmt"
	"manager/container"
	"manager/table"
	"os"

	"github.com/jmoiron/sqlx"
)

type DumpModel struct {
	Services []container.ContainerFull `json:"services"`
	Tables   []table.Table             `json:"tables"`
}

// Dump is export json file that service info and table info
func Dump(path string, db *sqlx.DB) error {
	// Get service data.
	services := fetchAllServices()
	// Get table data.
	tables := fetchAllTables(db)
	// Merge
	dump := DumpModel{
		Services: services,
		Tables:   tables,
	}

	// json encode
	data, err := json.Marshal(dump)
	if err != nil {
		return err
	}

	// Write to file.
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// dump all contaier info.
func fetchAllServices() []container.ContainerFull {
	cons, err := container.FetchAllServicesFull()
	if err != nil {
		fmt.Println(err)
		return []container.ContainerFull{}
	}
	
	return cons
}

// dump all table info.
func fetchAllTables(db *sqlx.DB) []table.Table {
	var tables []table.Table
	tbs := table.FetchAllTable(db)
	for _, tb := range tbs {
		t := table.FetchTableDetail(db, tb.TableName)
		tables = append(tables, t)
	}

	return tables
}