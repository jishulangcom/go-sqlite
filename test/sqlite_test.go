package test

import (
	"github.com/jishulangcom/go-sqlite"
	"testing"
)

func Test(t *testing.T) {
	sqlite.NewDB("test.db")

	sql := "CREATE TABLE table_name(column1 datatype);"
	sqlite.DB.Exec(sql)
}
