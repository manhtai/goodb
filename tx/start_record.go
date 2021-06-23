package tx

import (
	"goodb/constant"
	"goodb/file"
	"goodb/log"
)

type StartRecord struct {
	txNum int
}

func NewStartRecord(page *file.Page) *StartRecord {
	txNum := page.GetInt(constant.INT_SIZE)
	return &StartRecord{txNum: txNum}
}

func (r *StartRecord) op() int {
	return START
}

func (r *StartRecord) txNumber() int {
	return r.txNum
}

func (r *StartRecord) undo(tx *Transaction) {
}

func writeSTARTToLog(logMgr *log.LogManager, txNum int) int {
	record := make([]byte, 2*constant.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, START)
	page.SetInt(constant.INT_SIZE, txNum)
	return logMgr.Append(record)
}
