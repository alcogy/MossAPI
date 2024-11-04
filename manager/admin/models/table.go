package models

import (
	"manager/database/mysql"

	"github.com/jmoiron/sqlx"
)



func GetAllTables(db *sqlx.DB) []string {
	return mysql.FetchAllTable(db)
}

func GetTableDetail(db *sqlx.DB, table string) mysql.Table {
	return mysql.FetchTableDetail(db, table)
}

func CreateTable(db *sqlx.DB, table mysql.Table) error {
	return mysql.CreateTable(db, table)
}

func DeleteTableDetail(db *sqlx.DB, table string) error {
	return mysql.DeleteTable(db, table)
}
