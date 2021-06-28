package basic

import (
	"goodb/metadata"
	"goodb/parse"
	"goodb/plan"
	"goodb/tx"
)

type BasicUpdatePlanner struct {
	metadataMgr *metadata.MetadataManager
}

func NewBasicUpdatePlanner(metadataMgr *metadata.MetadataManager) plan.UpdatePlanner {
	return &BasicUpdatePlanner{
		metadataMgr: metadataMgr,
	}
}

func (planner *BasicUpdatePlanner) ExecuteInsert(stmt parse.InsertStatement, tx *tx.Transaction) int {
	p := plan.NewTablePlan(tx, stmt.TableName, planner.metadataMgr)
	modifyScan := p.OpenToUpdate()

	modifyScan.Insert()
	for i, fieldName := range stmt.Fields {
		val := stmt.Values[i]
		modifyScan.SetVal(fieldName, val)
	}
	modifyScan.Close()
	return 1
}

func (planner *BasicUpdatePlanner) ExecuteUpdate(stmt parse.UpdateStatement, tx *tx.Transaction) int {
	tablePlan := plan.NewTablePlan(tx, stmt.TableName, planner.metadataMgr)
	modifyPlan := plan.NewModifyPlan(tablePlan, stmt.Predicate)
	modifyScan := modifyPlan.OpenToUpdate()

	count := 0
	for modifyScan.Next() {
		val := stmt.Expression.Eval(modifyScan)
		modifyScan.SetVal(stmt.FieldName, val)
		count++
	}

	modifyScan.Close()
	return count
}

func (planner *BasicUpdatePlanner) ExecuteDelete(stmt parse.DeleteStatement, tx *tx.Transaction) int {
	tablePlan := plan.NewTablePlan(tx, stmt.TableName, planner.metadataMgr)
	modifyPlan := plan.NewModifyPlan(tablePlan, stmt.Predicate)
	modifyScan := modifyPlan.OpenToUpdate()

	count := 0
	for modifyScan.Next() {
		modifyScan.Delete()
		count++
	}

	modifyScan.Close()
	return count
}

func (planner *BasicUpdatePlanner) ExecuteCreateTable(stmt parse.CreateTableStatement, tx *tx.Transaction) int {
	planner.metadataMgr.CreateTable(stmt.TableName, &stmt.Schema, tx)
	return 0
}

func (planner *BasicUpdatePlanner) ExecuteCreateIndex(stmt parse.CreateIndexStatement, tx *tx.Transaction) int {
	planner.metadataMgr.CreateIndex(stmt.IndexName, stmt.TableName, stmt.FieldName, tx)
	return 0
}
