package tx

import (
	"goodb/constant"
	"goodb/file"
	"goodb/log"
)

type CommitRecord struct {
	txNum int
}

func NewCommitRecord(page *file.Page) *CommitRecord {
	txNum := page.GetInt(constant.INT_SIZE)
	return &CommitRecord{txNum: txNum}
}

func (r *CommitRecord) op() int {
	return COMMIT
}

func (r *CommitRecord) txNumber() int {
	return r.txNum
}

func (r *CommitRecord) undo(tx *Transaction) {
}

func writeCOMMITToLog(logMgr *log.LogManager, txNum int) int {
	record := make([]byte, 2*constant.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, COMMIT)
	page.SetInt(constant.INT_SIZE, txNum)
	return logMgr.Append(record)
}
