package parse

import (
	"goodb/query"
	"goodb/record"
)

type AstKind uint

const (
	SelectKind AstKind = iota
	CreateTableKind
	CreateIndexKind
	CreateViewKind
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

type CreateViewStatement struct {
	ViewName   string
	SelectStmt *SelectStatement
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
	Fields     []string
	Expression *query.Expression
	Predicate  *query.Predicate
}

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	CreateIndexStatement *CreateIndexStatement
	CreateViewStatement  *CreateViewStatement
	InsertStatement      *InsertStatement
	DeleteStatement      *DeleteStatement
	UpdateStatement      *UpdateStatement
	Kind                 AstKind
}

type Ast struct {
	Statements []*Statement
}
