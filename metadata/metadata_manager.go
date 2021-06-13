package metadata

import "goodb/tx"

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
