package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Table struct {
	Name        string `json:"name"`
	Columns 	[]Column `json:"columns"`
}

type Column struct {
	Name 	string `json:"name"`
	Type 	string `json:"type"`
	PK 		bool 	 `json:"pk"`
	Index bool 	 `json:"index"`
}

type TableInfo struct {
	Field string `db:"Field"`
	Type 	string `db:"Type"`
	Null 	string `db:"Null"`
	Key 	string `db:"Key"`
	Default 	string `db:"Default"`
	Extra 	string `db:"Extra"`
}

func Connection() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_HOST"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}

	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FetchAllTable(db *sqlx.DB) []string {
	var tables []string
	rows, err := db.Query(`show tables`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var tb string
		rows.Scan(&tb)
		tables = append(tables, tb)
	}

	return tables
}

func FetchTableDetail(db *sqlx.DB, tb string) Table {
	sql := fmt.Sprintf("SHOW columns FROM %s", tb)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return Table{}
	}

	var columns []Column
	for rows.Next() {
		var info TableInfo
		rows.Scan(&info.Field, &info.Type, &info.Null, &info.Key, &info.Default, &info.Extra)
		col := Column{
			Name: info.Field,
			Type: info.Type,
			PK: info.Key == "PRI",
			Index: info.Key == "MUL",
		}
		columns = append(columns, col)
	}
	
	table := Table{
		Name: tb,
		Columns: columns,
	}
	
	return table
}

func ExecuteDDL(db *sqlx.DB, ddl string) {
	_, err := db.Exec(ddl)
	if err != nil {
		panic(err)
	}
}

func DeleteTable(db *sqlx.DB, tb string) error {
	sql := fmt.Sprintf("drop table %s", tb)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}