package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)


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

func ExecuteDDL(db *sqlx.DB, ddl string) {
	_, err := db.Exec(ddl)
	if err != nil {
		panic(err)
	}
}