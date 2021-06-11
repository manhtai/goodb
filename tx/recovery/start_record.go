package recovery

import (
	"goodb/file"
	"goodb/log"
	"goodb/tx"
)

type StartRecord struct {
	txNum int
}

func NewStartRecord(page *file.Page) *StartRecord {
	txNum := page.GetInt(log.INT_SIZE)
	return &StartRecord{txNum: txNum}
}

func (r *StartRecord) op() int {
	return START
}

func (r *StartRecord) txNumber() int {
	return r.txNum
}

func (r *StartRecord) undo(tx *tx.Transaction) {
}

func writeSTARTToLog(logMgr *log.LogManager, txNum int) int {
	record := make([]byte, 2*log.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, START)
	page.SetInt(log.INT_SIZE, txNum)
	return logMgr.Append(record)
}
