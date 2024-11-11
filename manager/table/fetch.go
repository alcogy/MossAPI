package table

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// FetchAllTable returns all table info.
func FetchAllTable(db *sqlx.DB) []Table {
	var tables []Table
	rows, err := db.Queryx(`SHOW TABLE STATUS`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var ti TableInfo
		rows.StructScan(&ti)
		tb := Table{
			TableName: ti.Name,
			TableDesc: ti.Comment,
		}
		tables = append(tables, tb)
	}

	return tables
}

// FetchTableDetail returns column info by table.
func FetchTableDetail(db *sqlx.DB, tb string) Table {
	var table Table

	tableInfo := getTableInfo(db, tb)
	table.TableName = tableInfo.Name
	table.TableDesc = tableInfo.Comment
	
	columnsInfo := getColumnInfo(db, tb)
	indexesInfo := getIndexInfo(db, tb)

	var columns []Column
	for _, c := range columnsInfo {
		// Basic info.
		col := Column{
			Name: c.Field,
			Type: c.Type,
			PK: c.Key == "PRI",
			Nullable: c.Null != "NO",
			Comment: c.Comment,
		}

		// Index and unique info.
		for _, i := range indexesInfo {
			if i.ColumnName == c.Field {
				keys := strings.Split(i.KeyName, "_")
				// unique or index
				kind := keys[0]
				// index number
				num := keys[len(keys) - 1]
				if kind == "index" {
					val, err := strconv.Atoi(num)
					if err != nil {
						val = 0
					}
					col.Index = val
				} else if kind == "unique" {
					val, err := strconv.Atoi(num)
					if err != nil {
						val = 0
					}
					col.Unique = val
				}
			}
		}
		columns = append(columns, col)
	}
	
	table.Columns = columns

	return table
}

func getTableInfo(db *sqlx.DB, tb string) TableInfo {
	var tableInfo TableInfo
	sql := fmt.Sprintf("SHOW TABLE STATUS WHERE NAME = '%s'", tb)
	err := db.QueryRowx(sql).StructScan(&tableInfo)
	if err != nil {
		fmt.Println(err)
		return TableInfo{}
	}
	
	return tableInfo
}

func getColumnInfo(db *sqlx.DB, tb string) []ColumnInfo {
	sql := fmt.Sprintf("SHOW FULL COLUMNS FROM `%s`", tb)
	rows, err := db.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		return []ColumnInfo{}
	}
	var columns []ColumnInfo
	for rows.Next() {
		var ci ColumnInfo
		rows.StructScan(&ci)		
		columns = append(columns, ci)
	}

	return columns
}

func getIndexInfo(db *sqlx.DB, tb string) []IndexInfo {
	sql := fmt.Sprintf("SHOW index FROM `%s`", tb)
	rows, err := db.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		return []IndexInfo{}
	}

	var indexes []IndexInfo
	for rows.Next() {
		var ii IndexInfo
		rows.StructScan(&ii)		
		indexes = append(indexes, ii)
	}

	return indexes
}