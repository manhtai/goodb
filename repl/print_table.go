package repl

import (
	"fmt"
	"goodb/plan"
	"goodb/record"
	"io"
	"text/tabwriter"
)

func printTable(plan plan.Plan, out io.Writer) {
	scan := plan.Open()
	schema := plan.Schema()

	w := tabwriter.NewWriter(out, 1, 4, 0, ' ', tabwriter.Debug)
	for _, fldName := range schema.Fields() {
		fmt.Fprintf(w, "%s\t", fldName)
	}
	fmt.Fprintln(w)

	for range schema.Fields() {
		fmt.Fprint(w, "----\t")
	}
	fmt.Fprintln(w)

	for scan.Next() {
		for _, fldName := range schema.Fields() {
			if schema.Type(fldName) == record.INTEGER {
				fmt.Fprintf(w, "%d\t", scan.GetInt(fldName))
			} else {
				fmt.Fprintf(w, "%s\t", scan.GetString(fldName))
			}
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}
