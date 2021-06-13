package metadata

import (
	"goodb/record"
	"goodb/tx"
)

type IndexManager struct {
	layout   *record.Layout
	tableMgr *TableManager
	statMgr  *StatManager
}

func NewIndexManager(isNew bool, tableMgr *TableManager, statMgr *StatManager, tx *tx.Transaction) *IndexManager {
	if isNew {
		schema := &record.Schema{}
		schema.AddStringField("indexName", MAX_NAME)
		schema.AddStringField("tableName", MAX_NAME)
		schema.AddStringField("fieldName", MAX_NAME)
		tableMgr.createTable("idxCat", schema, tx)
	}

	layout := tableMgr.getLayout("idxCat", tx)
	return &IndexManager{
		layout:   layout,
		tableMgr: tableMgr,
		statMgr:  statMgr,
	}
}
