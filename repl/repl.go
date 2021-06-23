package repl

import (
	"bufio"
	"fmt"
	"goodb/parse"
	"goodb/server"
	"io"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var db *server.GooDb

	for {
		fmt.Printf("Choose a database name to connect: ")
		scanned := scanner.Scan()
		if !scanned {
			continue
		}
		dbDir := scanner.Text()
		db = server.NewGooDb(dbDir)
		break
	}

	for {
		fmt.Println()
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		parser := parse.NewParser(line)
		stmt := parser.ParseStatement()

		tx := db.NewTx()
		planner := db.Planner()

		switch stmt.Kind {
		case parse.SelectKind:
			plan := planner.CreateQueryPlan(stmt.SelectStatement, tx)
			printTable(plan, out)
		default:
			planner.ExecuteUpdatePlan(stmt, tx)
			tx.Commit()
		}
	}
}
