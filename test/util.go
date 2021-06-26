package test

import (
	"fmt"
	"goodb/parse"
	"goodb/record"
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
		schema := plan.Schema()

		count := 0
		for scan.Next() {
			for _, fldName := range schema.Fields() {
				if schema.Type(fldName) == record.INTEGER {
					scan.GetInt(fldName)
				} else {
					scan.GetString(fldName)
				}
			}
			count++
		}
		return count

	default:
		count := planner.ExecuteUpdatePlan(stmt, tx)
		tx.Commit()
		return count
	}
}
