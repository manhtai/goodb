package plan

import (
	"goodb/parse"
	"goodb/tx"
)

type QueryPlanner interface {
	CreatePlan(stmt parse.SelectStatement, tx *tx.Transaction) Plan
}

type UpdatePlaner interface {
	ExecuteInsert(stmt parse.InsertStatement, tx *tx.Transaction) int
	ExecuteUpdate(stmt parse.UpdateStatement, tx *tx.Transaction) int
	ExecuteDelete(stmt parse.DeleteStatement, tx *tx.Transaction) int

	ExecuteCreateTable(stmt parse.CreateTableStatement, tx *tx.Transaction) int
	ExecuteCreateIndex(stmt parse.CreateIndexStatement, tx *tx.Transaction) int
}

type Planner struct {
	queryPlanner  QueryPlanner
	updatePlanner UpdatePlaner
}

func NewPlanner(query QueryPlanner, update UpdatePlaner) *Planner {
	return &Planner{
		queryPlanner:  query,
		updatePlanner: update,
	}
}

func (planner *Planner) CreateQueryPlan(stmt parse.SelectStatement, tx *tx.Transaction) Plan {
	return planner.queryPlanner.CreatePlan(stmt, tx)
}

func (planner *Planner) ExecuteUpdatePlan(stmt parse.Statement, tx *tx.Transaction) int {
	switch stmt.Kind {
	case parse.InsertKind:
		return planner.updatePlanner.ExecuteInsert(stmt.InsertStatement, tx)
	case parse.DeleteKind:
		return planner.updatePlanner.ExecuteDelete(stmt.DeleteStatement, tx)
	case parse.UpdateKind:
		return planner.updatePlanner.ExecuteUpdate(stmt.UpdateStatement, tx)
	case parse.CreateTableKind:
		return planner.updatePlanner.ExecuteCreateTable(stmt.CreateTableStatement, tx)
	case parse.CreateIndexKind:
		return planner.updatePlanner.ExecuteCreateIndex(stmt.CreateIndexStatement, tx)
	default:
		return 0
	}
}
