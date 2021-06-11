package recovery

import (
	"goodb/file"
	"goodb/log"
	"goodb/tx"
)

type SetIntRecord struct {
	txNum int
	offset int
	val int
	block *file.Block
}

func NewSetIntRecord(page *file.Page) *SetIntRecord {
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
	val := page.GetInt(pos)

	return &SetIntRecord{
		txNum: txNum,
		offset: offset,
		val: val,
		block: block,
	}
}

func (r *SetIntRecord) op() int {
	return SETINT
}

func (r *SetIntRecord) txNumber() int {
	return r.txNum
}

func (r *SetIntRecord) undo(tx *tx.Transaction) {
	tx.Pin(r.block)
	tx.SetInt(r.block, r.offset, r.val, false)
	tx.Unpin(r.block)
}