package test

import (
	"fmt"
	"goodb/parse"
	"goodb/server"
)

func printStruct(v interface{}) string {
	return fmt.Sprintf("%+v", v)
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
