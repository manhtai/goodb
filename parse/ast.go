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
	ModifyKind
)

type SelectStatement struct {
	fields    []string
	tables    []string
	predicate *query.Predicate
}

type CreateTableStatement struct {
	tableName string
	schema    *record.Schema
}

type CreateViewStatement struct {
	viewName        string
	selectStatement *SelectStatement
}

type CreateIndexStatement struct {
	indexName string
	tableName string
	fieldName string
}

type InsertStatement struct {
	tableName string
	fields    []string
	values    []*query.Constant
}

type DeleteStatement struct {
	tableName string
	predicate *query.Predicate
}

type ModifyStatement struct {
	tableName  string
	fields     []string
	expression *query.Expression
	predicate  *query.Predicate
}

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	CreateIndexStatement *CreateIndexStatement
	CreateViewStatement  *CreateViewStatement
	InsertStatement      *InsertStatement
	DeleteStatement      *DeleteStatement
	ModifyStatement      *ModifyStatement
	Kind                 AstKind
}

type Ast struct {
	Statements []*Statement
}
