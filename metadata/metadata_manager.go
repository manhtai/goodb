package metadata

import (
	"goodb/record"
	"goodb/tx"
)

type MetadataManager struct {
	tableMgr *TableManager
	indexMgr *IndexManager
}

func NewMetadataManager(isNew bool, tx *tx.Transaction) *MetadataManager {
	tableMgr := NewTableManager(isNew, tx)
	indexMgr := NewIndexManager(isNew, tableMgr, tx)

	return &MetadataManager{
		tableMgr: tableMgr,
		indexMgr: indexMgr,
	}
}

func (metaMgr *MetadataManager) CreateTable(tblName string, schema *record.Schema, tx *tx.Transaction) {
	metaMgr.tableMgr.CreateTable(tblName, schema, tx)
}

func (metaMgr *MetadataManager) GetLayout(tblName string, tx *tx.Transaction) *record.Layout {
	return metaMgr.tableMgr.GetLayout(tblName, tx)
}

func (metaMgr *MetadataManager) CreateIndex(idxName string, tblName string, fldName string, tx *tx.Transaction) {
	metaMgr.indexMgr.createIndex(idxName, tblName, fldName, tx)
}

func (metaMgr *MetadataManager) GetIndexInfo(tblName string, tx *tx.Transaction) map[string]*IndexInfo {
	return metaMgr.indexMgr.getIndexInfo(tblName, tx)
}
