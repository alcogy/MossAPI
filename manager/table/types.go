package table

import "database/sql"

type Table struct {
	TableName string   `json:"tableName"`
	TableDesc string   `json:"tableDesc"`
	Columns   []Column `json:"columns"`
}

type Column struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	PK       bool   `json:"pk"`
	Nullable bool   `json:"nullable"`
	Unique   int    `json:"unique"`
	Index    int    `json:"index"`
	Comment  string `json:"comment"`
}

type TableInfo struct {
	Name          string         `db:"Name"`
	Engine        string         `db:"Engine"`
	Version       int32          `db:"Version"`
	RowFormat     string         `db:"Row_format"`
	Rows          int32          `db:"Rows"`
	AvgRowLength  int32          `db:"Avg_row_length"`
	DataLength    int32          `db:"Data_length"`
	MaxDataLength int32          `db:"Max_data_length"`
	IndexLength   int32          `db:"Index_length"`
	DataFree      int32          `db:"Data_free"`
	AutoIncrement sql.NullInt64  `db:"Auto_increment"`
	CreateTime    sql.NullTime   `db:"Create_time"`
	UpdateTime    sql.NullTime   `db:"Update_time"`
	CheckTime     sql.NullTime   `db:"Check_time"`
	Collation     string         `db:"Collation"`
	Checksum      sql.NullString `db:"Checksum"`
	CreateOptions string         `db:"Create_options"`
	Comment       string         `db:"Comment"`
}

type ColumnInfo struct {
	Field      string         `db:"Field"`
	Type       string         `db:"Type"`
	Collation  sql.NullString `db:"Collation"`
	Null       string         `db:"Null"`
	Key        string         `db:"Key"`
	Default    sql.NullString `db:"Default"`
	Extra      string         `db:"Extra"`
	Privileges string         `db:"Privileges"`
	Comment    string         `db:"Comment"`
}

type IndexInfo struct {
	Table        string         `db:"Table"`
	NonUnique    int            `db:"Non_unique"`
	KeyName      string         `db:"Key_name"`
	SeqInIndex   int            `db:"Seq_in_index"`
	ColumnName   string         `db:"Column_name"`
	Collation    sql.NullString `db:"Collation"`
	Cardinality  int            `db:"Cardinality"`
	SubPart      sql.NullString `db:"Sub_part"`
	Packed       sql.NullString `db:"Packed"`
	Null         string         `db:"Null"`
	IndexType    string         `db:"Index_type"`
	Comment      string         `db:"Comment"`
	IndexComment string         `db:"Index_comment"`
	Visible      string         `db:"Visible"`
	Expression   sql.NullString `db:"Expression"`
}
