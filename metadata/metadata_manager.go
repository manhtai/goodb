package metadata

import (
	"goodb/record"
	"goodb/tx"
)

type MetadataManager struct {
	tableMgr *TableManager
	viewMgr  *ViewManager
	statMgr  *StatManager
	indexMgr *IndexManager
}

func NewMetadataManager(isNew bool, tx *tx.Transaction) *MetadataManager {
	tableMgr := NewTableManager(isNew, tx)
	viewMgr := NewViewManager(isNew, tableMgr, tx)
	statMgr := NewStatManager(tableMgr, tx)
	indexMgr := NewIndexManager(isNew, tableMgr, statMgr, tx)

	return &MetadataManager{
		tableMgr: tableMgr,
		viewMgr:  viewMgr,
		statMgr:  statMgr,
		indexMgr: indexMgr,
	}
}

func (metaMgr *MetadataManager) CreateTable(tblName string, schema *record.Schema, tx *tx.Transaction) {
	metaMgr.tableMgr.createTable(tblName, schema, tx)
}

func (metaMgr *MetadataManager) GetLayout(tblName string, tx *tx.Transaction) *record.Layout {
	return metaMgr.tableMgr.getLayout(tblName, tx)
}

func (metaMgr *MetadataManager) CreateView(viewName string, viewDef string, tx *tx.Transaction) {
	metaMgr.viewMgr.createView(viewName, viewDef, tx)
}

func (metaMgr *MetadataManager) GetViewDef(viewName string, tx *tx.Transaction) string {
	return metaMgr.GetViewDef(viewName, tx)
}

func (metaMgr *MetadataManager) CreateIndex(idxName string, tblName string, fldName string, tx *tx.Transaction) {
	metaMgr.indexMgr.createIndex(idxName, tblName, fldName, tx)
}

func (metaMgr *MetadataManager) GetIndexInfo(tblName string, tx *tx.Transaction) map[string]*IndexInfo {
	return metaMgr.indexMgr.getIndexInfo(tblName, tx)
}

func (metaMgr *MetadataManager) GetStatInfo(tblName string, layout *record.Layout, tx *tx.Transaction) *StatInfo {
	return metaMgr.statMgr.getStatInfo(tblName, layout, tx)
}
