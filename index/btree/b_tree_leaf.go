package btree

import (
	"goodb/file"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type BTreeLeaf struct {
	tx          *tx.Transaction
	layout      record.Layout
	key         query.Constant
	data        BTreePage
	currentSlot int
	filename    string
}

func NewBLeaf(tx *tx.Transaction, block file.Block, layout record.Layout, key query.Constant) BTreeLeaf {
	data := NewBTreePage(tx, block, layout)
	return BTreeLeaf{
		tx:          tx,
		layout:      layout,
		key:         key,
		data:        data,
		filename:    block.Filename(),
		currentSlot: data.FindSlotBefore(key),
	}
}

func (l *BTreeLeaf) Close() {
	l.data.Close()
}

func (l *BTreeLeaf) Next() bool {
	l.currentSlot++
	if l.currentSlot >= l.data.GetNumRecs() {
		return l.tryOverflow()
	}

	if l.data.GetDataVal(l.currentSlot) == l.key {
		return true
	}

	return l.tryOverflow()
}

func (l *BTreeLeaf) GetDataRecord() record.Record {
	return l.data.GetDataRecord(l.currentSlot)
}

func (l *BTreeLeaf) Insert(rcd record.Record) *DirEntry {
	if l.data.GetFlag() >= 0 && l.data.GetDataVal(0).CompareTo(l.key) > 0 {
		firstVal := l.data.GetDataVal(0)
		block := l.data.Split(0, l.data.GetFlag())
		l.currentSlot = 0
		l.data.SetFlag(-1)
		l.data.InsertLeaf(l.currentSlot, l.key, rcd)
		return NewDirEntry(firstVal, block.Number())
	}

	l.currentSlot++
	l.data.InsertLeaf(l.currentSlot, l.key, rcd)
	if !l.data.IsFull() {
		return nil
	}

	firstKey := l.data.GetDataVal(0)
	lastKey := l.data.GetDataVal(l.data.GetNumRecs() - 1)

	if lastKey == firstKey {
		block := l.data.Split(1, l.data.GetFlag())
		l.data.SetFlag(block.Number())
		return nil
	}

	splitPos := l.data.GetNumRecs() / 2
	splitKey := l.data.GetDataVal(splitPos)

	if splitKey == firstKey {
		for l.data.GetDataVal(splitPos) == splitKey {
			splitPos++
		}
		splitKey = l.data.GetDataVal(splitPos)
	} else {
		for l.data.GetDataVal(splitPos-1) == splitKey {
			splitPos--
		}
	}
	block := l.data.Split(splitPos, -1)
	return NewDirEntry(splitKey, block.Number())
}

func (l *BTreeLeaf) Delete(rcd record.Record) {
	for l.Next() {
		if l.GetDataRecord() == rcd {
			l.data.Delete(l.currentSlot)
		}
	}
}

func (l *BTreeLeaf) tryOverflow() bool {
	firstKey := l.data.GetDataVal(0)
	flag := l.data.GetFlag()
	if l.key != firstKey || flag < 0 {
		return false
	}

	l.data.Close()
	nextBlock := file.NewBlock(l.filename, flag)
	l.data = NewBTreePage(l.tx, nextBlock, l.layout)
	l.currentSlot = 0
	return true
}
