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

func (tableScan *TableScan) BeforeFirst() {
	tableScan.moveToBlock(0)
}

func (tableScan *TableScan) Next() bool {
	tableScan.currentSlot = tableScan.recordPage.NextAfter(tableScan.currentSlot)
	for tableScan.currentSlot < 0 {
		if tableScan.atLastBlock() {
			return false
		}
		tableScan.moveToBlock(tableScan.recordPage.Block().Number() + 1)
		tableScan.currentSlot = tableScan.recordPage.NextAfter(tableScan.currentSlot)
	}
	return true
}

func (tableScan *TableScan) GetInt(fieldName string) int {
	return tableScan.recordPage.GetInt(tableScan.currentSlot, fieldName)
}

func (tableScan *TableScan) GetString(fieldName string) string {
	return tableScan.recordPage.GetString(tableScan.currentSlot, fieldName)
}

func (tableScan *TableScan) GetVal(fieldName string) *Constant {
	if tableScan.layout.Schema().Type(fieldName) == record.INTEGER {
		intVal := tableScan.recordPage.GetInt(tableScan.currentSlot, fieldName)
		return &Constant{intVal: intVal}
	}
	strVal := tableScan.recordPage.GetString(tableScan.currentSlot, fieldName)
	return &Constant{strVal: strVal}
}

func (tableScan *TableScan) HasField(fieldName string) bool {
	return tableScan.layout.Schema().HasField(fieldName)
}

func (tableScan *TableScan) Close() {
	if tableScan.recordPage != nil {
		tableScan.tx.Unpin(tableScan.recordPage.Block())
	}
}

func (tableScan *TableScan) SetInt(fieldName string, val int) {
	tableScan.recordPage.SetInt(tableScan.currentSlot, fieldName, val)
}

func (tableScan *TableScan) SetString(fieldName string, val string) {
	tableScan.recordPage.SetString(tableScan.currentSlot, fieldName, val)
}

func (tableScan *TableScan) SetVal(fieldName string, val *Constant) {
	if val.kind == StringKind {
		tableScan.SetString(fieldName, val.strVal)
	} else {
		tableScan.SetInt(fieldName, val.intVal)
	}
}

func (tableScan *TableScan) Insert() {
	tableScan.currentSlot = tableScan.recordPage.InsertAfter(tableScan.currentSlot)
	for tableScan.currentSlot < 0 {
		if tableScan.atLastBlock() {
			tableScan.moveToNewBlock()
		} else {
			tableScan.moveToBlock(tableScan.recordPage.Block().Number() + 1)
		}
		tableScan.currentSlot = tableScan.recordPage.InsertAfter(tableScan.currentSlot)
	}
}

func (tableScan *TableScan) Delete() {
	tableScan.recordPage.Delete(tableScan.currentSlot)
}

func (tableScan *TableScan) GetRecord() *record.Record {
	return record.NewRecord(
		tableScan.recordPage.Block().Number(),
		tableScan.currentSlot,
	)
}

func (tableScan *TableScan) MoveToRecord(rcd *record.Record) {
	tableScan.Close()
	block := file.NewBlock(tableScan.filename, rcd.BlockNumber())

	tableScan.recordPage = record.NewRecordPage(tableScan.tx, block, tableScan.layout)
	tableScan.currentSlot = rcd.Slot()
}

func (tableScan *TableScan) moveToNewBlock() {
	tableScan.Close()
	block := tableScan.tx.Append(tableScan.filename)
	tableScan.recordPage = record.NewRecordPage(tableScan.tx, block, tableScan.layout)
	tableScan.recordPage.Format()
	tableScan.currentSlot = -1
}

func (tableScan *TableScan) moveToBlock(blockNum int) {
	tableScan.Close()
	block := file.NewBlock(tableScan.filename, blockNum)
	tableScan.recordPage = record.NewRecordPage(tableScan.tx, block, tableScan.layout)
	tableScan.currentSlot = -1
}

func (tableScan *TableScan) atLastBlock() bool {
	return tableScan.recordPage.Block().Number() == tableScan.tx.Size(tableScan.filename)-1
}
