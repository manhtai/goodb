package btree

import (
	"fmt"
	"goodb/constant"
	"goodb/file"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type BTreeIndex struct {
	tx         *tx.Transaction
	dirLayout  record.Layout
	leafLayout record.Layout
	leafTable  string
	leaf       BTreeLeaf
	rootBlock  file.Block
}

func NewBTreeIndex(tx *tx.Transaction, idxName string, idxLayout record.Layout) BTreeIndex {
	// Leaf
	leafTbl := fmt.Sprintf("%sLeaf", idxName)
	if tx.Size(leafTbl) == 0 {
		block := tx.Append(leafTbl)
		page := NewBTreePage(tx, block, idxLayout)
		page.Format(block, -1)
	}

	// Dir
	schema := record.NewSchema()
	schema.AddSchema("currentBlock", *idxLayout.Schema())
	schema.AddSchema("dataVal", *idxLayout.Schema())

	dirTable := fmt.Sprintf("%sDir", idxName)
	dirLayout := record.NewLayoutFromSchema(*schema)
	rootBlock := file.NewBlock(dirTable, 0)

	if tx.Size(dirTable) == 0 {
		tx.Append(dirTable)
		page := NewBTreePage(tx, rootBlock, dirLayout)
		page.Format(rootBlock, -1)

		fldType := schema.Type("dataVal")
		minVal := query.NewIntConstant(constant.MIN_INT)
		if fldType == record.INTEGER {
			minVal = query.NewStrConstant("")
		}
		page.InsertDir(0, minVal, 0)
		page.Close()
	}

	return BTreeIndex{
		tx:         tx,
		leafLayout: idxLayout,
		leafTable:  leafTbl,
		rootBlock:  rootBlock,
	}
}

func (b *BTreeIndex) BeforeFirst(data query.Constant) {
	b.Close()
	root := NewBTreeDir(b.tx, b.rootBlock, b.dirLayout)
	blockNum := root.Search(data)
	root.Close()

	leafBlock := file.NewBlock(b.leafTable, blockNum)
	b.leaf = NewBLeaf(b.tx, leafBlock, b.leafLayout, data)
}

func (b BTreeIndex) Next() bool {
	return b.leaf.Next()
}

func (b BTreeIndex) GetRecord() record.Record {
	return b.leaf.GetDataRecord()
}

func (b BTreeIndex) Insert(data query.Constant, rcd record.Record) {
	b.BeforeFirst(data)
	entry := b.leaf.Insert(rcd)
	b.leaf.Close()
	if entry == nil {
		return
	}

	root := NewBTreeDir(b.tx, b.rootBlock, b.dirLayout)
	entry2 := root.Insert(entry)

	if entry2 != nil {
		root.MakeNewRoot(entry2)
	}
	root.Close()
}

func (b BTreeIndex) Delete(data query.Constant, rcd record.Record) {
	b.BeforeFirst(data)
	b.leaf.Delete(rcd)
	b.leaf.Close()
}

func (b BTreeIndex) Close() {
	b.leaf.Close()
}
