package btree

import (
	"goodb/constant"
	"goodb/file"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type BTreePage struct {
	tx           *tx.Transaction
	currentBlock file.Block
	layout       record.Layout
}

func NewBTreePage(tx *tx.Transaction, block file.Block, layout record.Layout) BTreePage {
	tx.Pin(block)
	return BTreePage{
		tx:           tx,
		currentBlock: block,
		layout:       layout,
	}
}

func (page BTreePage) FindSlotBefore(key query.Constant) int {
	slot := 0
	for slot < page.GetNumRecs() && page.GetDataVal(slot).CompareTo(key) < 0 {
		slot++
	}
	return slot - 1
}

func (page BTreePage) Close() {
	if page.currentBlock.Filename() != "" {
		page.tx.Unpin(page.currentBlock)
	}
}

func (page BTreePage) Format(block file.Block, flag int) {
	page.tx.SetInt(block, 0, flag, false)
	page.tx.SetInt(block, constant.INT_SIZE, 0, false)
	recordSize := page.layout.SlotSize()
	for pos := 2 * constant.INT_SIZE; pos+recordSize <= page.tx.BlockSize(); pos += recordSize {
		makeDefaultRecord(block, pos)
	}
}

func (page BTreePage) InsertDir(i int, val query.Constant, i2 int) {

}

func (page BTreePage) GetNumRecs() int {
	return page.tx.GetInt(page.currentBlock, constant.INT_SIZE)
}

func (page BTreePage) GetDataVal(slot int) query.Constant {
	return page.getVal(slot, "dataVal")
}

func (page BTreePage) GetDataRecord(slot int) record.Record {
	return record.NewRecord(page.getInt(slot, "block"), page.getInt(slot, "id"))
}

func (page BTreePage) Delete(slot int) {
	for i := slot + 1; i < page.GetNumRecs(); i++ {
		page.copyRecord(i, i-1)
		page.setNumRecs(page.GetNumRecs() - 1)
	}
}

func (page BTreePage) GetFlag() int {
	return page.tx.GetInt(page.currentBlock, 0)
}

func (page BTreePage) Split(splitPos int, flag int) file.Block {
	newBlock := page.appendNew(flag)
	newPage := NewBTreePage(page.tx, newBlock, page.layout)
	page.transferRecs(splitPos, newPage)
	newPage.SetFlag(flag)
	newPage.Close()
	return newBlock
}

func (page BTreePage) SetFlag(i int) {
	page.tx.SetInt(page.currentBlock, 0, i, true)
}

func (page BTreePage) InsertLeaf(slot int, key query.Constant, rcd record.Record) {
	page.insert(slot)
	page.setVal(slot, "dataVal", key)
	page.setInt(slot, "block", rcd.BlockNumber())
	page.setInt(slot, "id", rcd.Slot())
}

func (page BTreePage) IsFull() bool {
	return page.slotPos(page.GetNumRecs()+1) >= page.tx.BlockSize()
}

func (page BTreePage) GetChildNum(slot int) int {
	return page.getInt(slot, "currentBlock")
}

func (page BTreePage) appendNew(flag int) file.Block {
	block := page.tx.Append(page.currentBlock.Filename())
	page.tx.Pin(block)
	page.Format(block, flag)
	return block
}

func (page BTreePage) setInt(slot int, fldName string, val int) {
	pos := page.fldPos(slot, fldName)
	page.tx.SetInt(page.currentBlock, pos, val, true)
}

func (page BTreePage) setString(slot int, fldName string, val string) {
	pos := page.fldPos(slot, fldName)
	page.tx.SetString(page.currentBlock, pos, val, true)
}

func (page BTreePage) setVal(slot int, fldName string, val query.Constant) {
	schema := page.layout.Schema()
	if schema.Type(fldName) == record.INTEGER {
		page.setInt(slot, fldName, val.Int())
	}
	page.setString(slot, fldName, val.Str())
}
func (page BTreePage) getInt(slot int, fldName string) int {
	pos := page.fldPos(slot, fldName)
	return page.tx.GetInt(page.currentBlock, pos)
}

func (page BTreePage) getString(slot int, fldName string) string {
	pos := page.fldPos(slot, fldName)
	return page.tx.GetString(page.currentBlock, pos)
}

func (page BTreePage) getVal(slot int, fldName string) query.Constant {
	schema := page.layout.Schema()
	if schema.Type(fldName) == record.INTEGER {
		return query.NewIntConstant(page.getInt(slot, fldName))
	}
	return query.NewStrConstant(page.getString(slot, fldName))
}

func (page BTreePage) setNumRecs(n int) {
	page.tx.SetInt(page.currentBlock, constant.INT_SIZE, n, true)
}

func (page BTreePage) insert(slot int) {
	for i := page.GetNumRecs(); i > slot; i-- {
		page.copyRecord(i-1, i)
		page.setNumRecs(page.GetNumRecs() + 1)
	}
}

func (page BTreePage) copyRecord(from int, to int) {
	schema := page.layout.Schema()
	for _, fldName := range schema.Fields() {
		page.setVal(to, fldName, page.getVal(from, fldName))
	}

}

func (page BTreePage) transferRecs(slot int, dest BTreePage) {
	destSlot := 0
	for slot < page.GetNumRecs() {
		dest.insert(destSlot)
		schema := page.layout.Schema()
		for _, fldName := range schema.Fields() {
			dest.setVal(destSlot, fldName, page.getVal(slot, fldName))
		}
		page.Delete(slot)
		destSlot++
	}
}

func (page BTreePage) fldPos(slot int, fldName string) int {
	offset := page.layout.Offset(fldName)
	return page.slotPos(slot) + offset
}

func (page BTreePage) slotPos(slot int) int {
	slotSize := page.layout.SlotSize()
	return constant.INT_SIZE*2 + (slot * slotSize)
}

func makeDefaultRecord(block file.Block, pos int) {

}
