package btree

import (
	"goodb/file"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type DirEntry struct {
	val      query.Constant
	blockNum int
}

func NewDirEntry(val query.Constant, number int) *DirEntry {
	return &DirEntry{
		val:      val,
		blockNum: number,
	}
}

type BTreeDir struct {
	tx       *tx.Transaction
	layout   record.Layout
	data     BTreePage
	filename string
}

func NewBTreeDir(tx *tx.Transaction, block file.Block, layout record.Layout) BTreeDir {
	return BTreeDir{
		tx:       tx,
		layout:   layout,
		data:     NewBTreePage(tx, block, layout),
		filename: block.Filename(),
	}
}

func (dir BTreeDir) Search(key query.Constant) int {
	childBlock := dir.findChildBlock(key)
	for dir.data.GetFlag() > 0 {
		dir.data.Close()
		dir.data = NewBTreePage(dir.tx, childBlock, dir.layout)
		childBlock = dir.findChildBlock(key)
	}
	return childBlock.Number()
}

func (dir BTreeDir) Close() {
	dir.data.Close()
}

func (dir BTreeDir) Insert(entry *DirEntry) *DirEntry {
	if dir.data.GetFlag() == 0 {
		return dir.insertEntry(entry)
	}

	childBlock := dir.findChildBlock(entry.val)
	child := NewBTreeDir(dir.tx, childBlock, dir.layout)
	myEntry := dir.Insert(entry)
	child.Close()

	if myEntry != nil {
		dir.insertEntry(myEntry)
	}
	return nil
}

func (dir BTreeDir) MakeNewRoot(entry *DirEntry) {
	firstVal := dir.data.GetDataVal(0)
	level := dir.data.GetFlag()
	block := dir.data.Split(0, level)
	oldRoot := NewDirEntry(firstVal, block.Number())
	dir.insertEntry(oldRoot)
	dir.insertEntry(entry)
	dir.data.SetFlag(level + 1)
}

func (dir BTreeDir) findChildBlock(key query.Constant) file.Block {
	slot := dir.data.FindSlotBefore(key)
	if dir.data.GetDataVal(slot+1) == key {
		slot++
	}
	blockNum := dir.data.GetChildNum(slot)
	return file.NewBlock(dir.filename, blockNum)
}

func (dir BTreeDir) insertEntry(entry *DirEntry) *DirEntry {
	newSlot := dir.data.FindSlotBefore(entry.val)
	dir.data.InsertDir(newSlot, entry.val, entry.blockNum)
	if !dir.data.IsFull() {
		return nil
	}

	level := dir.data.GetFlag()
	splitPos := dir.data.GetNumRecs() / 2
	splitVal := dir.data.GetDataVal(splitPos)
	block := dir.data.Split(splitPos, level)
	return NewDirEntry(splitVal, block.Number())
}
