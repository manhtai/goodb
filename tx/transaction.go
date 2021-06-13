package tx

import (
	"fmt"
	"goodb/buffer"
	"goodb/file"
	"goodb/tx/concurrency"
	"goodb/tx/recovery"
)

const END_OF_FILE = -1

type Transaction struct {
	recoveryMgr *recovery.RecoveryManager
	concurrencyMgr *concurrency.ConcurrencyManager
	bufferMgr *buffer.BufferManager
	fileMgr *file.FileManager
	buffers *BufferList
	nextTxNum int
	txNum int
}

func (tx *Transaction) Commit() {
	tx.recoveryMgr.Commit();
	fmt.Printf("Transaction %d committed", tx.txNum)
	tx.concurrencyMgr.Release()
	tx.buffers.unpinAll()
}

func (tx *Transaction) Rollback() {
	tx.recoveryMgr.Rollback();
	fmt.Printf("Transaction %d rolled back", tx.txNum)
	tx.concurrencyMgr.Release()
	tx.buffers.unpinAll()
}

func (tx *Transaction) Recover() {
	tx.bufferMgr.FlushAll(tx.txNum)
	tx.recoveryMgr.Recover()
}

func (tx *Transaction) Pin(block *file.Block) {
	tx.buffers.pin(block)
}

func (tx *Transaction) Unpin(block *file.Block) {
	tx.buffers.unpin(block)
}

func (tx *Transaction) GetInt(block *file.Block, offset int) int {
	tx.concurrencyMgr.SLock(block)
	buff := tx.buffers.getBuffer(block)
	return buff.Contents().GetInt(offset)
}
func (tx *Transaction) GetString(block *file.Block, offset int) string {
	tx.concurrencyMgr.SLock(block)
	buff := tx.buffers.getBuffer(block)
	return buff.Contents().GetString(offset)
}

func (tx *Transaction) SetInt(block *file.Block, offset int, val int, okToLog bool) {
	tx.concurrencyMgr.XLock(block)
	buff := tx.buffers.getBuffer(block)

	lsn := -1
	if okToLog {
		lsn = tx.recoveryMgr.SetInt(buff, offset)
	}

	page := buff.Contents()
	page.SetInt(offset, val)
	buff.SetModified(tx.txNum, lsn)
}

func (tx *Transaction) SetString(block *file.Block, offset int, val string, okToLog bool) {
	tx.concurrencyMgr.XLock(block)
	buff := tx.buffers.getBuffer(block)

	lsn := -1
	if okToLog {
		lsn = tx.recoveryMgr.SetString(buff, offset)
	}

	page := buff.Contents()
	page.SetString(offset, val)
	buff.SetModified(tx.txNum, lsn)
}

func (tx *Transaction) Size(filename string) int {
	dummyBlock := file.NewBlock(filename, END_OF_FILE)
	tx.concurrencyMgr.SLock(dummyBlock)
	return tx.fileMgr.Length(filename)
}

func (tx *Transaction) Append(filename string) *file.Block {
	dummyBlock := file.NewBlock(filename, END_OF_FILE)
	tx.concurrencyMgr.XLock(dummyBlock)
	return tx.fileMgr.Append(filename)
}

func (tx *Transaction) BlockSize() int {
	return tx.fileMgr.BlockSize()
}