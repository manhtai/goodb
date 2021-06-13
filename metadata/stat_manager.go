package metadata

import "goodb/tx"

type StatManager struct {
	tableMgr  *TableManager
	tableStat map[string]*StatInfo
	numCalls  int
}

func NewStatManager(tableMgr *TableManager, tx *tx.Transaction) *StatManager {
	statMgr := &StatManager{tableMgr: tableMgr}
	statMgr.refreshStats(tx)
	return statMgr
}

func (statMgr *StatManager) refreshStats(tx *tx.Transaction) {
	panic("implement me")
}
