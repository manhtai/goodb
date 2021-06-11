package tx

import (
	"goodb/buffer"
	"goodb/file"
	"goodb/tx/concurrency"
	"goodb/tx/recovery"
)

type Transaction struct {
	recoveryMgr *recovery.RecoveryManager
	concurrencyMgr *concurrency.ConcurrencyManager
	bufferMgr *buffer.BufferManager
	fileMgr *file.FileManager
	buffers *BufferList
	nextTxNum int
	END_OF_FILE int
	txNum int
}

func (tx *Transaction) Pin(block *file.Block) {

}

func (tx *Transaction) Unpin(block *file.Block) {

}

func (tx *Transaction) SetInt(block *file.Block, offset int, val int, okToLog bool) {

}

func (tx *Transaction) SetString(block *file.Block, offset int, val string, okToLog bool) {
}