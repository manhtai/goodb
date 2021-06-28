package test

import (
	"goodb/file"
	"goodb/server"
	"os"
	"testing"
)

func TestOptCRUD(t *testing.T) {
	os.RemoveAll(file.DB_DIR_PREFIX)
	db := server.NewGooDbOpt("test")

	// CREATE TABLE
	parseAndExecute(db, "create table test(i int, v varchar(10))")
	parseAndExecute(db, "create table test_case(ti int, name varchar(10))")

	// CREATE INDEX
	parseAndExecute(db, "create index idx_test_i on test(i)")
	parseAndExecute(db, "create index idx_test_v on test(v)")
}
