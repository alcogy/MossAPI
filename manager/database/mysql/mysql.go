package mysql

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

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