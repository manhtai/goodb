package tx

import (
	"goodb/constant"
	"goodb/file"
	"goodb/log"
)

type SetIntRecord struct {
	txNum  int
	offset int
	val    int
	block  file.Block
}

func NewSetIntRecord(page *file.Page) *SetIntRecord {
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
	val := page.GetInt(pos)

	return &SetIntRecord{
		txNum:  txNum,
		offset: offset,
		val:    val,
		block:  block,
	}
}

func (r *SetIntRecord) op() int {
	return SETINT
}

func (r *SetIntRecord) txNumber() int {
	return r.txNum
}

func (r *SetIntRecord) undo(tx *Transaction) {
	tx.Pin(r.block)
	tx.SetInt(r.block, r.offset, r.val, false)
	tx.Unpin(r.block)
}

func WriteSETINToLog(logMgr *log.LogManager, txNum int, block file.Block, offset int, val int) int {
	txPos := constant.INT_SIZE
	filePos := txPos + constant.INT_SIZE
	blockPos := filePos + log.MaxLength(len(block.Filename()))
	offsetPos := blockPos + constant.INT_SIZE
	valuePos := offsetPos + constant.INT_SIZE

	record := make([]byte, valuePos+constant.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, SETINT)
	page.SetInt(txPos, txNum)
	page.SetString(filePos, block.Filename())
	page.SetInt(blockPos, block.Number())
	page.SetInt(offsetPos, offset)
	page.SetInt(valuePos, val)

	return logMgr.Append(record)
}
