package plan

import (
	"goodb/parse"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type Plan interface {
	Open() query.Scan
	Schema() record.Schema
}

type UpdatePlan interface {
	Plan

	OpenToUpdate() query.UpdateScan
}

type QueryPlanner interface {
	CreatePlan(stmt *parse.SelectStatement, tx *tx.Transaction) Plan
}

type UpdatePlaner interface {
	ExecuteInsert(stmt *parse.InsertStatement, tx *tx.Transaction) int
	ExecuteUpdate(stmt *parse.UpdateStatement, tx *tx.Transaction) int
	ExecuteDelete(stmt *parse.DeleteStatement, tx *tx.Transaction) int

	ExecuteCreateTable(stmt *parse.CreateTableStatement, tx *tx.Transaction) int
	ExecuteCreateView(smtm *parse.CreateViewStatement, tx *tx.Transaction) int
	ExecuteCreateIndex(stmt *parse.CreateIndexStatement, tx *tx.Transaction) int
}
