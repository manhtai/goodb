package opt

import (
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
