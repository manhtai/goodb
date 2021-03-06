package plan

import (
	"goodb/metadata"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type TablePlan struct {
	tableName string
	tx        *tx.Transaction
	layout    record.Layout
}

func NewTablePlan(tx *tx.Transaction, tableName string, metadataMgr *metadata.MetadataManager) UpdatePlan {
	layout := metadataMgr.GetLayout(tableName, tx)
	return TablePlan{
		tableName: tableName,
		tx:        tx,
		layout:    layout,
	}
}

func (tp TablePlan) OpenToUpdate() query.UpdateScan {
	return query.NewTableScan(tp.tx, tp.tableName, tp.layout)
}

func (tp TablePlan) Open() query.Scan {
	return query.NewTableScan(tp.tx, tp.tableName, tp.layout)
}

func (tp TablePlan) Schema() record.Schema {
	return *tp.layout.Schema()
}
