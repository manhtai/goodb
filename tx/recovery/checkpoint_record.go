package recovery

import (
	"goodb/file"
	"goodb/log"
	"goodb/tx"
)

type CheckpointRecord struct {
}

func (r *CheckpointRecord) op() int {
	return CHECKPOINT
}

func (r *CheckpointRecord) txNumber() int {
	return -1
}

func (r *CheckpointRecord) undo(tx *tx.Transaction) {
}

func writeCHECKPOINTToLog(logMgr *log.LogManager) int {
	record := make([]byte, log.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, CHECKPOINT)
	return logMgr.Append(record)
}