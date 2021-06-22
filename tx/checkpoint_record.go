package tx

import (
	"goodb/constant"
	"goodb/file"
	"goodb/log"
)

type CheckpointRecord struct {
}

func (r *CheckpointRecord) op() int {
	return CHECKPOINT
}

func (r *CheckpointRecord) txNumber() int {
	return -1
}

func (r *CheckpointRecord) undo(tx *Transaction) {
}

func writeCHECKPOINTToLog(logMgr *log.LogManager) int {
	record := make([]byte, constant.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, CHECKPOINT)
	return logMgr.Append(record)
}
