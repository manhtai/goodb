package plan

import (
	"goodb/metadata"
	"goodb/parse"
	"goodb/tx"
)

type QueryPlanner interface {
	CreatePlan(stmt parse.SelectStatement, tx *tx.Transaction) Plan
}

type UpdatePlanner interface {
	ExecuteInsert(stmt parse.InsertStatement, tx *tx.Transaction) int
	ExecuteUpdate(stmt parse.UpdateStatement, tx *tx.Transaction) int
	ExecuteDelete(stmt parse.DeleteStatement, tx *tx.Transaction) int

	ExecuteCreateTable(stmt parse.CreateTableStatement, tx *tx.Transaction) int
	ExecuteCreateIndex(stmt parse.CreateIndexStatement, tx *tx.Transaction) int
}

type QueryPlannerFunc func(metadataMgr *metadata.MetadataManager) QueryPlanner
type UpdatePlannerFunc func(metadataMgr *metadata.MetadataManager) UpdatePlanner

type Planner struct {
	queryPlanner  QueryPlanner
	updatePlanner UpdatePlanner
}

func NewPlanner(query QueryPlanner, update UpdatePlanner) *Planner {
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
