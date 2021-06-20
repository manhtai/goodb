package repl

import (
	"bufio"
	"fmt"
	"goodb/parse"
	"goodb/server"
	"io"
)

const PROMPT = ">>> "

// Start read text from stdin and print token to stdout
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var db *server.GooDb

	for {
		fmt.Printf("Choose a database name...")
		scanned := scanner.Scan()
		if !scanned {
			continue
		}
		dbDir := scanner.Text()
		db = server.NewGooDb(dbDir)
		break
	}

	tx := db.NewTx()
	planner := db.Planner();

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		parser := parse.NewParser(line)
		stmt := parser.ParseStatement()

		switch stmt.Kind {
		case parse.SelectKind:
			plan := planner.CreateQueryPlan(stmt.SelectStatement, tx)
			scan := plan.Open()
			for scan.Next() {
				// TODO: Display values in table format
			}
		default:
			planner.ExecuteUpdatePlan(stmt, tx)
		}
	}
}
