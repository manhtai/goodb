package query

import (
	"fmt"
	"goodb/file"
	"goodb/record"
	"goodb/tx"
)

type TableScan struct {
	tx          *tx.Transaction
	layout      *record.Layout
	recordPage  *record.RecordPage
	filename    string
	currentSlot int
}

func NewTableScan(tx *tx.Transaction, tableName string, layout *record.Layout) *TableScan {
	filename := fmt.Sprintf("%s.tbl", tableName)
	tableScan := &TableScan{
		tx:       tx,
		layout:   layout,
		filename: filename,
	}
	if tx.Size(filename) == 0 {
		tableScan.moveToNewBlock()
	} else {
		tableScan.moveToBlock(0)
	}
	return tableScan
}

func (tableScan *TableScan) beforeFirst() {
	tableScan.moveToBlock(0)
}


func (tableScan *TableScan) next() bool {
	rp := tableScan.recordPage
	for sl := rp.NextAfter(tableScan.currentSlot); sl < 0; sl = rp.NextAfter(sl) {
		if tableScan.atLastBlock() {
			return false
		}
		tableScan.moveToBlock(tableScan.recordPage.Block().Number() + 1)
	}
	return true
}

func (tableScan *TableScan) getInt(fieldName string) int {
	return tableScan.recordPage.GetInt(tableScan.currentSlot, fieldName)
}

func (tableScan *TableScan) getString(fieldName string) string {
	return tableScan.recordPage.GetString(tableScan.currentSlot, fieldName)
}

func (tableScan *TableScan) hasField(fieldName string) bool {
	return tableScan.layout.Schema().HasField(fieldName)
}

func (tableScan *TableScan) close() {
	if tableScan.recordPage != nil {
		tableScan.tx.Unpin(tableScan.recordPage.Block())
	}
}

func (tableScan *TableScan) setInt(fieldName string, val int) {
	tableScan.recordPage.SetInt(tableScan.currentSlot, fieldName, val)
}

func (tableScan *TableScan) setString(fieldName string, val string) {
	tableScan.recordPage.SetString(tableScan.currentSlot, fieldName, val)
}

func (tableScan *TableScan) insert() {
	rp := tableScan.recordPage
	for sl := rp.InsertAfter(tableScan.currentSlot) ; sl < 0; sl = rp.InsertAfter(sl) {
		if tableScan.atLastBlock() {
			tableScan.moveToNewBlock()
		} else {
			tableScan.moveToBlock(rp.Block().Number() + 1)
		}
	}
}

func (tableScan *TableScan) delete() {
	tableScan.recordPage.Delete(tableScan.currentSlot)
}

func (tableScan *TableScan) getRecord() *record.Record {
	return record.NewRecord(
		tableScan.recordPage.Block().Number(),
		tableScan.currentSlot,
	)
}

func (tableScan *TableScan) moveToRecord(rcrd *record.Record) {
	tableScan.close()
	block := file.NewBlock(tableScan.filename, rcrd.BlockNumber())

	tableScan.recordPage = record.NewRecordPage(tableScan.tx, block , tableScan.layout)
	tableScan.currentSlot = rcrd.Slot()
}

func (tableScan *TableScan) moveToNewBlock() {
	tableScan.close()
	block := tableScan.tx.Append(tableScan.filename)
	tableScan.recordPage = record.NewRecordPage(tableScan.tx, block, tableScan.layout)
	tableScan.recordPage.Format()
	tableScan.currentSlot = -1
}

func (tableScan *TableScan) moveToBlock(blockNum int) {
	tableScan.close()
	block := file.NewBlock(tableScan.filename, blockNum)
	tableScan.recordPage = record.NewRecordPage(tableScan.tx, block, tableScan.layout)
	tableScan.currentSlot = -1
}

func (tableScan *TableScan) atLastBlock() bool {
	return tableScan.recordPage.Block().Number() == tableScan.tx.Size(tableScan.filename)-1
}