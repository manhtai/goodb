package opt

import (
	"goodb/index"
	"goodb/metadata"
	"goodb/parse"
	"goodb/plan"
	"goodb/tx"
)

type IndexUpdatePlanner struct {
	metadataMgr *metadata.MetadataManager
}

func NewIndexUpdatePlanner(metadataMgr *metadata.MetadataManager) plan.UpdatePlanner {
	return &IndexUpdatePlanner{
		metadataMgr: metadataMgr,
	}
}

func (planner *IndexUpdatePlanner) ExecuteInsert(stmt parse.InsertStatement, tx *tx.Transaction) int {
	tblName := stmt.TableName
	p := plan.NewTablePlan(tx, tblName, planner.metadataMgr)

	updateScan := p.OpenToUpdate()
	updateScan.Insert()
	record := updateScan.GetRecord()

	indexes := planner.metadataMgr.GetIndexInfo(tblName, tx)
	for i, fldName := range stmt.Fields {
		val := stmt.Values[i]
		updateScan.SetVal(fldName, val)

		idxInfo := indexes[fldName]
		if idxInfo != nil {
			idx := idxInfo.Open()
			idx.Insert(val, record)
			idx.Close()
		}
	}

	updateScan.Close()
	return 1
}

func (planner *IndexUpdatePlanner) ExecuteUpdate(stmt parse.UpdateStatement, tx *tx.Transaction) int {
	tblName := stmt.TableName
	fldName := stmt.FieldName

	tablePlan := plan.NewTablePlan(tx, tblName, planner.metadataMgr)
	modifyPlan := plan.NewModifyPlan(tablePlan, stmt.Predicate)
	modifyScan := modifyPlan.OpenToUpdate()

	indexes := planner.metadataMgr.GetIndexInfo(tblName, tx)
	indexInfo := indexes[fldName]
	var idx index.Index
	if indexInfo != nil {
		idx = indexInfo.Open()
	}

	count := 0
	for modifyScan.Next() {
		oldVal := modifyScan.GetVal(fldName)
		val := stmt.Expression.Eval(modifyScan)
		modifyScan.SetVal(stmt.FieldName, val)
		count++

		if idx != nil {
			rcd := modifyScan.GetRecord()
			idx.Delete(oldVal, rcd)
			idx.Insert(val, rcd)
		}
	}

	modifyScan.Close()
	return count
}

func (planner *IndexUpdatePlanner) ExecuteDelete(stmt parse.DeleteStatement, tx *tx.Transaction) int {
	tblName := stmt.TableName
	tablePlan := plan.NewTablePlan(tx, tblName, planner.metadataMgr)
	modifyPlan := plan.NewModifyPlan(tablePlan, stmt.Predicate)
	modifyScan := modifyPlan.OpenToUpdate()

	indexes := planner.metadataMgr.GetIndexInfo(tblName, tx)

	count := 0
	for modifyScan.Next() {
		record := modifyScan.GetRecord()
		for fldName := range indexes {
			val := modifyScan.GetVal(fldName)
			idx := indexes[fldName].Open()
			idx.Delete(val, record)
			idx.Close()
		}
		modifyScan.Delete()
		count++
	}
	return count
}

func (planner *IndexUpdatePlanner) ExecuteCreateTable(stmt parse.CreateTableStatement, tx *tx.Transaction) int {
	planner.metadataMgr.CreateTable(stmt.TableName, stmt.Schema, tx)
	return 0
}

func (planner *IndexUpdatePlanner) ExecuteCreateIndex(stmt parse.CreateIndexStatement, tx *tx.Transaction) int {
	planner.metadataMgr.CreateIndex(stmt.IndexName, stmt.TableName, stmt.FieldName, tx)
	return 0
}
