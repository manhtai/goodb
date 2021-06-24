package test

import (
	"goodb/parse"
	"goodb/query"
	"goodb/record"
	"testing"
)

func TestCreateTableStatement(t *testing.T) {
	tests := []struct {
		input  string
		output parse.CreateTableStatement
	}{
		{"create table a (b int, c varchar(50))", parse.CreateTableStatement{
			TableName: "a",
			Schema:    *record.NewSchema().AddIntField("b").AddStringField("c", 50),
		}},
		{"create table a (b int, c varchar(50), d int)", parse.CreateTableStatement{
			TableName: "a",
			Schema:    *record.NewSchema().AddIntField("b").AddStringField("c", 50).AddIntField("d"),
		}},
	}

	for _, inout := range tests {
		parser := parse.NewParser(inout.input)
		stmt := parser.ParseStatement()
		if stmt.Kind != parse.CreateTableKind || printStruct(stmt.CreateTableStatement) != printStruct(inout.output) {
			t.Errorf("Expect:\n %+v\nGot:\n %+v", inout.output, stmt.CreateTableStatement)
		}
	}
}

func TestInsertStatement(t *testing.T) {
	tests := []struct {
		input  string
		output parse.InsertStatement
	}{
		{"insert into a(b) values ('2')", parse.InsertStatement{
			TableName: "a",
			Fields:    []string{"b"},
			Values: []query.Constant{
				query.NewStrConstant("2"),
			},
		}},
		{"insert into a(b, c) values (1, '2')", parse.InsertStatement{
			TableName: "a",
			Fields:    []string{"b", "c"},
			Values: []query.Constant{
				query.NewIntConstant(1),
				query.NewStrConstant("2"),
			},
		}},
		{"insert into a(b, c, d) values (1, '2', 'bar')", parse.InsertStatement{
			TableName: "a",
			Fields:    []string{"b", "c", "d"},
			Values: []query.Constant{
				query.NewIntConstant(1),
				query.NewStrConstant("2"),
				query.NewStrConstant("bar"),
			},
		}},
	}

	for _, inout := range tests {
		parser := parse.NewParser(inout.input)
		stmt := parser.ParseStatement()
		if stmt.Kind != parse.InsertKind || printStruct(stmt.InsertStatement) != printStruct(inout.output) {
			t.Errorf("Expect:\n %+v\nGot:\n %+v", inout.output, stmt.InsertStatement)
		}
	}
}

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

func TestUpdateStatement(t *testing.T) {
	tests := []struct {
		input  string
		output parse.UpdateStatement
	}{
		{"update a set b = 1", parse.UpdateStatement{
			TableName:  "a",
			FieldName:  "b",
			Expression: query.NewConstantExpression(query.NewIntConstant(1)),
		}},
		{"update a set b = c", parse.UpdateStatement{
			TableName:  "a",
			FieldName:  "b",
			Expression: query.NewFieldExpression("c"),
		}},
		{"update a set b = c where a = 1", parse.UpdateStatement{
			TableName:  "a",
			FieldName:  "b",
			Expression: query.NewFieldExpression("c"),
			Predicate: query.NewPredicateFromTerms([]query.Term{
				query.NewTerm(
					query.NewFieldExpression("a"),
					query.NewConstantExpression(query.NewIntConstant(1)),
				),
			}),
		}},
		{"update a set b = c where a = 1 and b = d", parse.UpdateStatement{
			TableName:  "a",
			FieldName:  "b",
			Expression: query.NewFieldExpression("c"),
			Predicate: query.NewPredicateFromTerms([]query.Term{
				query.NewTerm(
					query.NewFieldExpression("a"),
					query.NewConstantExpression(query.NewIntConstant(1)),
				),
				query.NewTerm(
					query.NewFieldExpression("b"),
					query.NewFieldExpression("d"),
				),
			}),
		}},
	}

	for _, inout := range tests {
		parser := parse.NewParser(inout.input)
		stmt := parser.ParseStatement()
		if stmt.Kind != parse.UpdateKind || printStruct(stmt.UpdateStatement) != printStruct(inout.output) {
			t.Errorf("Expect:\n %+v\nGot:\n %+v", inout.output, stmt.UpdateStatement)
		}
	}
}

func TestDeleteStatement(t *testing.T) {
	tests := []struct {
		input  string
		output parse.DeleteStatement
	}{
		{"delete from b where a = 1", parse.DeleteStatement{
			TableName: "b",
			Predicate: query.NewPredicateFromTerms([]query.Term{
				query.NewTerm(
					query.NewFieldExpression("a"),
					query.NewConstantExpression(query.NewIntConstant(1)),
				),
			}),
		}},
		{"delete from c where a = 1 and b = '2'", parse.DeleteStatement{
			TableName: "c",
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
		{"delete from d where a = 1 and b = '2' and 'bar' = c", parse.DeleteStatement{
			TableName: "d",
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
		if stmt.Kind != parse.DeleteKind || printStruct(stmt.DeleteStatement) != printStruct(inout.output) {
			t.Errorf("Expect:\n %+v\nGot:\n %+v", inout.output, stmt.DeleteStatement)
		}
	}
}
