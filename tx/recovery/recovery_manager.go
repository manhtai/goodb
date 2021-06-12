package recovery

import (
	"goodb/buffer"
	"goodb/log"
	"goodb/tx"
)

type RecoveryManager struct {
	logMgr    *log.LogManager
	bufferMgr *buffer.BufferManager
	tx        *tx.Transaction
	txNum     int
}

func NewRecoveryManager(
	tx *tx.Transaction,
	txNum int,
	logMgr *log.LogManager,
	bufferMgr *buffer.BufferManager,
) *RecoveryManager {
	writeSTARTToLog(logMgr, txNum)
	return &RecoveryManager{
		logMgr: logMgr,
		bufferMgr: bufferMgr,
		tx: tx,
		txNum: txNum,
	}
}

func (recoveryMgr *RecoveryManager) Commit() {
	recoveryMgr.bufferMgr.FlushAll(recoveryMgr.txNum)
	lsn := writeCOMMITToLog(recoveryMgr.logMgr, recoveryMgr.txNum)
	recoveryMgr.logMgr.Flush(lsn)
}

func (recoveryMgr *RecoveryManager) Rollback() {
	recoveryMgr.doRollback()
	recoveryMgr.bufferMgr.FlushAll(recoveryMgr.txNum)
	lsn := writeROLLBACKToLog(recoveryMgr.logMgr, recoveryMgr.txNum)
	recoveryMgr.logMgr.Flush(lsn)
}

func (recoveryMgr *RecoveryManager) Recover() {
	recoveryMgr.doRecover()
	recoveryMgr.bufferMgr.FlushAll(recoveryMgr.txNum)
	lsn := writeCHECKPOINTToLog(recoveryMgr.logMgr)
	recoveryMgr.logMgr.Flush(lsn)
}

func (recoveryMgr *RecoveryManager) SetInt(buff *buffer.Buffer, offset int) int {
	oldVal := buff.Contents().GetInt(offset)
	block := buff.Block()
	return WriteSETINToLog(recoveryMgr.logMgr, recoveryMgr.txNum, block, offset, oldVal)
}

func (recoveryMgr *RecoveryManager) SetString(buff *buffer.Buffer, offset int) int {
	oldVal := buff.Contents().GetString(offset)
	block := buff.Block()
	return WriteSETSTRINGoLog(recoveryMgr.logMgr, recoveryMgr.txNum, block, offset, oldVal)
}

func (recoveryMgr *RecoveryManager) doRollback() {
}

func (recoveryMgr *RecoveryManager) doRecover() {
}
