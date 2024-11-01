package models

import (
	"manager/database/mysql"

	"github.com/jmoiron/sqlx"
)

type Table struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

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