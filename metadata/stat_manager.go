package metadata

import (
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

const REFRESH_CALLS = 100

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

func (statMgr *StatManager) getStatInfo(tblName string, layout *record.Layout, tx *tx.Transaction) *StatInfo {
	statMgr.numCalls += 1
	if statMgr.numCalls > REFRESH_CALLS {
		statMgr.refreshStats(tx)
	}
	statInfo, ok := statMgr.tableStat[tblName]
	if !ok {
		statInfo = statMgr.getTableStats(tblName, layout, tx)
		statMgr.tableStat[tblName] = statInfo
	}
	return statInfo
}

func (statMgr *StatManager) refreshStats(tx *tx.Transaction) {
	statMgr.tableStat = make(map[string]*StatInfo)
	statMgr.numCalls = 0
	tblCatLayout := statMgr.tableMgr.getLayout("tblCat", tx)
	tblCatScan := query.NewTableScan(tx, "tblCat", tblCatLayout)

	for ; tblCatScan.Next(); {
		tblName := tblCatScan.GetString("tblName")
		layout := statMgr.tableMgr.getLayout("tblName", tx)
		statInfo := statMgr.getTableStats(tblName, layout, tx)
		statMgr.tableStat[tblName] = statInfo
	}
	tblCatScan.Close()
}

func (statMgr *StatManager) getTableStats(tblName string, layout *record.Layout, tx *tx.Transaction) *StatInfo {
	numRecords := 0
	numBlocks := 0
	tblScan := query.NewTableScan(tx, tblName, layout)
	for ; tblScan.Next() ; {
		numRecords += 1
		numBlocks = tblScan.GetRecord().BlockNumber() + 1
	}
	tblScan.Close()
	return &StatInfo{
		numRecords: numRecords,
		numBlocks: numBlocks,
	}
}