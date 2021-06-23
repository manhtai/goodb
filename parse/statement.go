package parse

import (
	"goodb/query"
	"goodb/record"
)

type StatementKind int

const (
	SelectKind StatementKind = iota
	CreateTableKind
	CreateIndexKind
	InsertKind
	DeleteKind
	UpdateKind
)

type SelectStatement struct {
	Fields    []string
	Tables    []string
	Predicate *query.Predicate
}

type CreateTableStatement struct {
	TableName string
	Schema    *record.Schema
}

type CreateIndexStatement struct {
	IndexName string
	TableName string
	FieldName string
}

type InsertStatement struct {
	TableName string
	Fields    []string
	Values    []*query.Constant
}

type DeleteStatement struct {
	TableName string
	Predicate *query.Predicate
}

type UpdateStatement struct {
	TableName  string
	FieldName  string
	Expression *query.Expression
	Predicate  *query.Predicate
}

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	CreateIndexStatement *CreateIndexStatement
	InsertStatement      *InsertStatement
	DeleteStatement      *DeleteStatement
	UpdateStatement      *UpdateStatement
	Kind                 StatementKind
}
