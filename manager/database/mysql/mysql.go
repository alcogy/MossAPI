package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Table struct {
	TableName     string `json:"tableName"`
	TableDesc     string `json:"tableDesc"`
	Columns 		[]Column `json:"columns"`
}

type Column struct {
	Name 	string `json:"name"`
	Type 	string `json:"type"`
	PK 		bool 	 `json:"pk"`
	Nullable bool `json:"nullable"`
	Unique int 	 `json:"unique"`
	Index int 	 `json:"index"`
	Comment string `json:"comment"`
}

type TableInfo struct {
	Name string `db:"Name"`
	Engine string `db:"Engine"`
	Version int32 `db:"Version"`
	RowFormat string `db:"Row_format"`
	Rows int32 `db:"Rows"`
	AvgRowLength int32 `db:"Avg_row_length"`
	DataLength int32 `db:"Data_length"`
	MaxDataLength int32 `db:"Max_data_length"`
	IndexLength int32 `db:"Index_length"`
	DataFree int32 `db:"Data_free"`
	AutoIncrement sql.NullInt64 `db:"Auto_increment"`
	CreateTime sql.NullTime `db:"Create_time"`
	UpdateTime sql.NullTime `db:"Update_time"`
	CheckTime sql.NullTime `db:"Check_time"`
	Collation string `db:"Collation"`
	Checksum sql.NullString `db:"Checksum"`
	CreateOptions string `db:"Create_options"`
	Comment string `db:"Comment"`
}

type ColumnInfo struct {
	Field string `db:"Field"`
	Type 	string `db:"Type"`
	Collation sql.NullString `db:"Collation"`
	Null 	string `db:"Null"`
	Key 	string `db:"Key"`
	Default 	sql.NullString `db:"Default"`
	Extra 	string `db:"Extra"`
	Privileges string `db:"Privileges"`
	Comment string `db:"Comment"`
}

type IndexInfo struct {
	Table string `db:"Table"`
	NonUnique int `db:"Non_unique"`
	KeyName string `db:"Key_name"`
	SeqInIndex int `db:"Seq_in_index"`
	ColumnName string `db:"Column_name"`
	Collation sql.NullString `db:"Collation"`
	Cardinality int `db:"Cardinality"`
	SubPart sql.NullString `db:"Sub_part"`
	Packed sql.NullString `db:"Packed"`
	Null string `db:"Null"`
	IndexType string `db:"Index_type"`
	Comment string `db:"Comment"`
	IndexComment string `db:"Index_comment"`
	Visible string `db:"Visible"`
	Expression sql.NullString `db:"Expression"`
}

// Connection To DDL
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
		ParseTime: true,
	}

	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ExecuteDDL(db *sqlx.DB, ddl string) {
	_, err := db.Exec(ddl)
	if err != nil {
		panic(err)
	}
}

func CreateTable(db *sqlx.DB, tb Table) error {
	sql := makeCreateTableSql(tb)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTable(db *sqlx.DB, tb string) error {
	sql := fmt.Sprintf("drop table %s", tb)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func makeCreateTableSql(table Table) string {
	var sql string
	var pks []string
	uniques := make(map[int][]string)
	indexes := make(map[int][]string)

	sql += fmt.Sprintf("create table `%s` (\n", table.TableName)
	for _, col := range table.Columns {
		notNull := ""
		if !col.Nullable {
			notNull = "not null"
		}

		sql += fmt.Sprintf("  `%s` %s %s comment '%s',\n", col.Name, col.Type, notNull, col.Comment)
		// Set PK columns.
		if col.PK {
			pks = append(pks, "`" + col.Name + "`")
		}

		// Set unique columns.
		if col.Unique > 0 {
			uniques[col.Unique] = append(uniques[col.Unique], "`" + col.Name + "`")
		}

		// Set Index columns.
		if col.Index > 0 {
			indexes[col.Index] = append(indexes[col.Index], "`" + col.Name + "`")
		}
	}
	
	if len(pks) > 0 {
		sql += fmt.Sprintf("  primary key (%s),\n", strings.Join(pks, ","))
	}

	for i, unique := range uniques {
		col := strings.Join(unique, ",")
		sql += fmt.Sprintf("  unique key `unique_%s_%d` (%s),\n", table.TableName, i, col)
	}

	for i, index := range indexes {
		col := strings.Join(index, ",")
		sql += fmt.Sprintf("  index `index_%s_%d` (%s),\n", table.TableName, i, col)
	}
	
	sql = sql[:len(sql)-2]
	sql += "\n)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='" + table.TableDesc + "';"

	return sql
}