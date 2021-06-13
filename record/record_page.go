package record

import (
	"goodb/file"
	"goodb/tx"
)

const (
	EMPTY = iota
	USED
)

type RecordPage struct {
	tx     *tx.Transaction
	block  *file.Block
	layout *Layout
}

func NewRecordPage(tx *tx.Transaction, block *file.Block, layout *Layout) *RecordPage {
	tx.Pin(block)
	return &RecordPage{
		tx:     tx,
		block:  block,
		layout: layout,
	}
}

func (rp *RecordPage) GetInt(slot int, fieldName string) int {
	fieldPos := rp.offset(slot) + rp.layout.Offset(fieldName)
	return rp.tx.GetInt(rp.block, fieldPos)
}

func (rp *RecordPage) GetString(slot int, fieldName string) string {
	fieldPos := rp.offset(slot) + rp.layout.Offset(fieldName)
	return rp.tx.GetString(rp.block, fieldPos)
}

func (rp *RecordPage) SetInt(slot int, fieldName string, val int) {
	fieldPos := rp.offset(slot) + rp.layout.Offset(fieldName)
	rp.tx.SetInt(rp.block, fieldPos, val, true)
}

func (rp *RecordPage) SetString(slot int, fieldName string, val string) {
	fieldPos := rp.offset(slot) + rp.layout.Offset(fieldName)
	rp.tx.SetString(rp.block, fieldPos, val, true)
}

func (rp *RecordPage) Delete(slot int) {
	rp.setFlag(slot, EMPTY)
}

func (rp *RecordPage) NextAfter(slot int) int {
	return rp.searchAfter(slot, USED)
}

func (rp *RecordPage) InsertAfter(slot int) int {
	newSlot := rp.searchAfter(slot, EMPTY)
	if newSlot >= 0 {
		rp.setFlag(slot, USED)
	}
	return newSlot
}

func (rp *RecordPage) searchAfter(slot int, flag int) int {
	slot += 1
	for ; rp.isValidSlot(slot); slot++ {
		if rp.tx.GetInt(rp.block, rp.offset(slot)) == flag {
			return slot
		}
	}
	return -1
}

func (rp *RecordPage) isValidSlot(slot int) bool {
	return rp.offset(slot) <= rp.tx.BlockSize()
}

func (rp *RecordPage) setFlag(slot int, flag int) {
	rp.tx.SetInt(rp.block, rp.offset(slot), flag, true)
}

func (rp *RecordPage) offset(slot int) int {
	return rp.layout.slotSize * slot
}
