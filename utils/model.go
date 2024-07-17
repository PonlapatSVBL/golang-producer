package utils

import (
	"database/sql"
	"fmt"

	"github.com/PonlapatSVBL/golang-producer/connections/mysql"
)

func ExportModel(table string) {
	var columns []struct {
		Field   string         `json:"Field" db:"Field"`
		Type    string         `json:"Type" db:"Type"`
		Null    string         `json:"Null" db:"Null"`
		Key     string         `json:"Key" db:"Key"`
		Default sql.NullString `json:"Default" db:"Default"`
		Extra   string         `json:"Extra" db:"Extra"`
	}

	query := fmt.Sprintf("SHOW COLUMNS FROM %s;", table)

	mysqlInstance := mysql.NewMysql()
	mysqlInstance.SqlList(&columns, query)

	for _, col := range columns {
		// fmt.Printf("Field: %s, Type: %s, Null: %s, Key: %s, Default: %s, Extra: %s\n", col.Field, col.Type, col.Null, col.Key, col.Default.String, col.Extra)
		fmt.Printf("%s string `json:\"%s\" db:\"%s\"`\n", ToCamelCase(col.Field), col.Field, col.Field)
	}
}
