package test

import (
	"goodb/file"
	"goodb/parse"
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

	// RETRIEVE
	c := parseAndExecute(db, "select i, v from test")
	if c != 2 {
		t.Errorf("Expect 2, got %d", c)
	}

	c = parseAndExecute(db, "select i, v from test where i = 1")
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

func parseAndExecute(db *server.GooDb, text string) int {
	tx := db.NewTx()
	planner := db.Planner()
	parser := parse.NewParser(text)
	stmt := parser.ParseStatement()

	switch stmt.Kind {
	case parse.SelectKind:
		plan := planner.CreateQueryPlan(stmt.SelectStatement, tx)
		scan := plan.Open()
		count := 0
		for scan.Next() {
			count++
		}
		return count

	default:
		count := planner.ExecuteUpdatePlan(stmt, tx)
		tx.Commit()
		return count
	}
}