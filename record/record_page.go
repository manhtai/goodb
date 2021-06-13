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

func (recordPage *RecordPage) Format() {
	slot := 0
	for ; recordPage.isValidSlot(slot); slot++ {
		recordPage.tx.SetInt(recordPage.block, recordPage.offset(slot), EMPTY, false)
		schema := recordPage.layout.Schema()
		for _, fieldName := range schema.Fields() {
			fieldPos := recordPage.offset(slot) + recordPage.layout.Offset(fieldName)
			if schema.Type(fieldName) == INTEGER {
				recordPage.tx.SetInt(recordPage.block, fieldPos, 0, false)
			} else {
				recordPage.tx.SetString(recordPage.block, fieldPos, "", false)
			}
		}
	}
}

func (recordPage *RecordPage) GetInt(slot int, fieldName string) int {
	fieldPos := recordPage.offset(slot) + recordPage.layout.Offset(fieldName)
	return recordPage.tx.GetInt(recordPage.block, fieldPos)
}

func (recordPage *RecordPage) GetString(slot int, fieldName string) string {
	fieldPos := recordPage.offset(slot) + recordPage.layout.Offset(fieldName)
	return recordPage.tx.GetString(recordPage.block, fieldPos)
}

func (recordPage *RecordPage) SetInt(slot int, fieldName string, val int) {
	fieldPos := recordPage.offset(slot) + recordPage.layout.Offset(fieldName)
	recordPage.tx.SetInt(recordPage.block, fieldPos, val, true)
}

func (recordPage *RecordPage) SetString(slot int, fieldName string, val string) {
	fieldPos := recordPage.offset(slot) + recordPage.layout.Offset(fieldName)
	recordPage.tx.SetString(recordPage.block, fieldPos, val, true)
}

func (recordPage *RecordPage) Delete(slot int) {
	recordPage.setFlag(slot, EMPTY)
}

func (recordPage *RecordPage) NextAfter(slot int) int {
	return recordPage.searchAfter(slot, USED)
}

func (recordPage *RecordPage) InsertAfter(slot int) int {
	newSlot := recordPage.searchAfter(slot, EMPTY)
	if newSlot >= 0 {
		recordPage.setFlag(slot, USED)
	}
	return newSlot
}

func (recordPage *RecordPage) searchAfter(slot int, flag int) int {
	slot += 1
	for ; recordPage.isValidSlot(slot); slot++ {
		if recordPage.tx.GetInt(recordPage.block, recordPage.offset(slot)) == flag {
			return slot
		}
	}
	return -1
}

func (recordPage *RecordPage) isValidSlot(slot int) bool {
	return recordPage.offset(slot) <= recordPage.tx.BlockSize()
}

func (recordPage *RecordPage) setFlag(slot int, flag int) {
	recordPage.tx.SetInt(recordPage.block, recordPage.offset(slot), flag, true)
}

func (recordPage *RecordPage) offset(slot int) int {
	return recordPage.layout.slotSize * slot
}
