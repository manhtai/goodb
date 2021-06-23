package tx

import (
	"goodb/constant"
	"goodb/file"
	"goodb/log"
)

type RollbackRecord struct {
	txNum int
}

func NewRollbackRecord(page *file.Page) *RollbackRecord {
	txNum := page.GetInt(constant.INT_SIZE)
	return &RollbackRecord{txNum: txNum}
}

func (r *RollbackRecord) op() int {
	return COMMIT
}

func (r *RollbackRecord) txNumber() int {
	return r.txNum
}

func (r *RollbackRecord) undo(tx *Transaction) {
}

func writeROLLBACKToLog(logMgr *log.LogManager, txNum int) int {
	record := make([]byte, 2*constant.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, ROLLBACK)
	page.SetInt(constant.INT_SIZE, txNum)
	return logMgr.Append(record)
}
