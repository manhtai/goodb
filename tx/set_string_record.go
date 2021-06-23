package tx

import (
	"goodb/constant"
	"goodb/file"
	"goodb/log"
)

type SetStringRecord struct {
	txNum  int
	offset int
	val    string
	block  file.Block
}

func NewSetStringRecord(page *file.Page) *SetStringRecord {
	pos := constant.INT_SIZE
	txNum := page.GetInt(pos)

	pos += constant.INT_SIZE
	filename := page.GetString(pos)

	pos += log.MaxLength(len(filename))
	blockNum := page.GetInt(pos)
	block := file.NewBlock(filename, blockNum)

	pos += constant.INT_SIZE
	offset := page.GetInt(pos)

	pos += constant.INT_SIZE
	val := page.GetString(pos)

	return &SetStringRecord{
		txNum:  txNum,
		offset: offset,
		val:    val,
		block:  block,
	}
}

func (r *SetStringRecord) op() int {
	return SETSTRING
}

func (r *SetStringRecord) txNumber() int {
	return r.txNum
}

func (r *SetStringRecord) undo(tx *Transaction) {
	tx.Pin(r.block)
	tx.SetString(r.block, r.offset, r.val, false)
	tx.Unpin(r.block)
}

func WriteSETSTRINGoLog(logMgr *log.LogManager, txNum int, block file.Block, offset int, val string) int {
	txPos := constant.INT_SIZE
	filePos := txPos + constant.INT_SIZE
	blockPos := filePos + log.MaxLength(len(block.Filename()))
	offsetPos := blockPos + constant.INT_SIZE
	valuePos := offsetPos + constant.INT_SIZE

	record := make([]byte, valuePos+log.MaxLength(len(val)))
	page := file.NewPageFromBytes(record)
	page.SetInt(0, SETINT)
	page.SetInt(txPos, txNum)
	page.SetString(filePos, block.Filename())
	page.SetInt(blockPos, block.Number())
	page.SetInt(offsetPos, offset)
	page.SetString(valuePos, val)

	return logMgr.Append(record)
}
