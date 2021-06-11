package recovery

import (
	"goodb/file"
	"goodb/log"
	"goodb/tx"
)

type SetStringRecord struct {
	txNum int
	offset int
	val string
	block *file.Block
}

func NewSetStringRecord(page *file.Page) *SetStringRecord {
	pos := log.INT_SIZE
	txNum := page.GetInt(pos)

	pos += log.INT_SIZE
	filename := page.GetString(pos)

	pos += len(filename)
	blockNum := page.GetInt(pos)
	block := file.NewBlock(filename, blockNum)

	pos += log.INT_SIZE
	offset := page.GetInt(pos)

	pos += log.INT_SIZE
	val := page.GetString(pos)

	return &SetStringRecord{
		txNum: txNum,
		offset: offset,
		val: val,
		block: block,
	}
}

func (r *SetStringRecord) op() int {
	return SETSTRING
}

func (r *SetStringRecord) txNumber() int {
	return r.txNum
}

func (r *SetStringRecord) undo(tx *tx.Transaction) {
	tx.Pin(r.block)
	tx.SetString(r.block, r.offset, r.val, false)
	tx.Unpin(r.block)
}
