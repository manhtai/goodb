package test

import (
	"fmt"
	"goodb/parse"
	"goodb/query"
	"testing"
)

func TestSelectStatement(t *testing.T) {
	tests := []struct {
		input  string
		output parse.SelectStatement
	}{
		{"select a from b ", parse.SelectStatement{
			Tables: []string{"b"},
			Fields: []string{"a"},
		}},
		{"select a, b from c", parse.SelectStatement{
			Tables: []string{"c"},
			Fields: []string{"a", "b"},
		}},
		{"select a, b from c where a = 1", parse.SelectStatement{
			Tables: []string{"c"},
			Fields: []string{"a", "b"},
			Predicate: query.NewPredicateFromTerms([]query.Term{
				query.NewTerm(
					query.NewFieldExpression("a"),
					query.NewConstantExpression(query.NewIntConstant(1)),
				),
			}),
		}},
		{"select a, b, c from d", parse.SelectStatement{
			Tables: []string{"d"},
			Fields: []string{"a", "b", "c"},
		}},
		{"select a, b, c from d where a = 1 and b = '2'", parse.SelectStatement{
			Tables: []string{"d"},
			Fields: []string{"a", "b", "c"},
			Predicate: query.NewPredicateFromTerms([]query.Term{
				query.NewTerm(
					query.NewFieldExpression("a"),
					query.NewConstantExpression(query.NewIntConstant(1)),
				),
				query.NewTerm(
					query.NewFieldExpression("b"),
					query.NewConstantExpression(query.NewStrConstant("2")),
				),
			}),
		}},
		{"select a, b, c from d where a = 1 and b = '2' and 'bar' = c", parse.SelectStatement{
			Tables: []string{"d"},
			Fields: []string{"a", "b", "c"},
			Predicate: query.NewPredicateFromTerms([]query.Term{
				query.NewTerm(
					query.NewFieldExpression("a"),
					query.NewConstantExpression(query.NewIntConstant(1)),
				),
				query.NewTerm(
					query.NewFieldExpression("b"),
					query.NewConstantExpression(query.NewStrConstant("2")),
				),
				query.NewTerm(
					query.NewConstantExpression(query.NewStrConstant("bar")),
					query.NewFieldExpression("c"),
				),
			}),
		}},
	}

	for _, inout := range tests {
		parser := parse.NewParser(inout.input)
		stmt := parser.ParseStatement()
		if stmt.Kind != parse.SelectKind || printStruct(stmt.SelectStatement) != printStruct(inout.output) {
			t.Errorf("Expect:\n %+v\nGot:\n %+v", inout.output, stmt.SelectStatement)
		}
	}
}

func printStruct(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}
