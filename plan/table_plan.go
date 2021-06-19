package plan

import (
	"goodb/metadata"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type TablePlan struct {
	tableName string
	tx *tx.Transaction
	layout *record.Layout
	stats *metadata.StatInfo
}

func NewTablePlan(tx *tx.Transaction, tableName string, metadataMgr *metadata.MetadataManager) *TablePlan {
	layout := metadataMgr.GetLayout(tableName, tx)
	stats := metadataMgr.GetStatInfo(tableName, layout, tx)
	return &TablePlan{
		tableName: tableName,
		tx: tx,
		layout: layout,
		stats: stats,
	}
}

func (tp TablePlan) Open() *query.TableScan {
	return query.NewTableScan(tp.tx, tp.tableName, tp.layout)
}

func (tp TablePlan) Schema() *record.Schema {
	return tp.layout.Schema()
}
