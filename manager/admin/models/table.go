package models

import (
	"manager/database/mysql"

	"github.com/jmoiron/sqlx"
)



func GetAllTables(db *sqlx.DB) []string {
	tables := mysql.FetchAllTable(db)
	return tables
	// return []Table{	
	// 	{
	// 		Name:        "product",
	// 		Description: "Management product.",
	// 	},
	// 	{
	// 		Name:        "customer_product",
	// 		Description: "Relation customer and product.",
	// 	},
	// }
}

func GetTableDetail(db *sqlx.DB, table string) mysql.Table {
	return mysql.FetchTableDetail(db, table)
}