package test

import (
	"goodb/file"
	"goodb/server"
	"os"
	"testing"
)

func TestCRUD(t *testing.T) {
	os.RemoveAll(file.DB_DIR_PREFIX)
	db := server.NewGooDb("test")

	// CREATE
	parseAndExecute(db, "create table test(i int, v varchar(10))")
	parseAndExecute(db, "insert into test(i, v) values (1, 'hi')")
	parseAndExecute(db, "insert into test(i, v) values (2, 'bar')")

	parseAndExecute(db, "create table test_case(ti int, name varchar(10))")
	parseAndExecute(db, "insert into test_case(ti, name) values (1, 'baz')")

	// RETRIEVE
	c := parseAndExecute(db, "select i, v from test")
	if c != 2 {
		t.Errorf("Expect 2, got %d", c)
	}

	c = parseAndExecute(db, "select i, v from test where i = 1")
	if c != 1 {
		t.Errorf("Expect 1, got %d", c)
	}

	c = parseAndExecute(db, "select i, v, ti, name from test, test_case")
	if c != 2 {
		t.Errorf("Expect 2, got %d", c)
	}

	c = parseAndExecute(db, "select i, v, ti, name from test, test_case where i = ti")
	if c != 1 {
		t.Errorf("Expect 1, got %d", c)
	}

	// UPDATE
	c = parseAndExecute(db, "update test set v = 'baz' where i = 1")
	if c != 1 {
		t.Errorf("Expect 1, got %d", c)
	}

	c = parseAndExecute(db, "select i, v from test where v = 'baz'")
	if c != 1 {
		t.Errorf("Expect 1, got %d", c)
	}

	// DELETE
	c = parseAndExecute(db, "delete from test where v = 'baz'")
	if c != 1 {
		t.Errorf("Expect 1, got %d", c)
	}

	c = parseAndExecute(db, "select i, v from test where v = 'baz'")
	if c != 0 {
		t.Errorf("Expect 0, got %d", c)
	}

	c = parseAndExecute(db, "select i, v from test where v = 'bar'")
	if c != 1 {
		t.Errorf("Expect 1, got %d", c)
	}
}
