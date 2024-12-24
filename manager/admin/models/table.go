package models

import (
	"manager/table"

	"github.com/jmoiron/sqlx"
)

func GetAllTables(db *sqlx.DB) []table.Table {
	return table.FetchAllTable(db)
}

func GetTableDetail(db *sqlx.DB, tb string) table.Table {
	return table.FetchTableDetail(db, tb)
}

func CreateTable(db *sqlx.DB, tb table.Table) error {
	return table.CreateTable(db, tb)
}

func DeleteTableDetail(db *sqlx.DB, tb string) error {
	return table.DeleteTable(db, tb)
}
