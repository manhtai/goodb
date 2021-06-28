package opt

import (
	"goodb/metadata"
	"goodb/parse"
	"goodb/plan"
	"goodb/tx"
)

type IndexUpdatePlanner struct {

}

func NewIndexUpdatePlanner(metadataMgr *metadata.MetadataManager) plan.UpdatePlanner {
	panic("Implement me!")
}

func (planner *IndexUpdatePlanner) ExecuteInsert(stmt parse.InsertStatement, tx *tx.Transaction) int {
	panic("Implement me!")
}

func (planner *IndexUpdatePlanner)ExecuteUpdate(stmt parse.UpdateStatement, tx *tx.Transaction) int {
	panic("Implement me!")
}

func (planner *IndexUpdatePlanner)ExecuteDelete(stmt parse.DeleteStatement, tx *tx.Transaction) int{
	panic("Implement me!")
}

func (planner *IndexUpdatePlanner)ExecuteCreateTable(stmt parse.CreateTableStatement, tx *tx.Transaction) int {
	panic("Implement me!")
}

func (planner *IndexUpdatePlanner)ExecuteCreateIndex(stmt parse.CreateIndexStatement, tx *tx.Transaction) int {
	panic("Implement me!")
}
