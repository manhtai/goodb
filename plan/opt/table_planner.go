package opt

import (
	"goodb/metadata"
	"goodb/plan"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type TablePlanner struct {
	tablePlan plan.UpdatePlan
	predicate query.Predicate
	schema record.Schema
	indexes map[string]*metadata.IndexInfo
	tx *tx.Transaction
}

func NewTablePlanner(tblName string, pred query.Predicate, tx *tx.Transaction, mgr *metadata.MetadataManager) *TablePlanner {
	tblPlan := plan.NewTablePlan(tx, tblName, mgr)
	return &TablePlanner{
		tablePlan: tblPlan,
		predicate: pred,
		schema: tblPlan.Schema(),
		indexes: mgr.GetIndexInfo(tblName, tx),
		tx: tx,
	}
}

func (tp *TablePlanner) MakeSelectPlan() {
	panic("Implement me!")
}

func (tp *TablePlanner) MakeJoinPlan() {
	panic("Implement me!")
}

func (tp *TablePlanner) MakeProductPlan() {
	panic("Implement me!")
}
