package btree

import (
	"goodb/file"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type BTreeIndex struct{
	tx *tx.Transaction
	dirLayout record.Layout
	leafLayout record.Layout
	leafTable string
	leaf BTreeLeaf
	rootBlock file.Block
}

func NewBTreeIndex(tx *tx.Transaction, idxName string, idxLayout record.Layout) BTreeIndex {
	return BTreeIndex{
		tx: tx,
	}
}

func (BTreeIndex) BeforeFirst(key query.Constant) {
	panic("implement me")
}

func (BTreeIndex) Next() bool {
	panic("implement me")
}

func (BTreeIndex) GetRecord() record.Record {
	panic("implement me")
}

func (BTreeIndex) Insert(data query.Constant, rcd record.Record) {
	panic("implement me")
}

func (BTreeIndex) Delete(data query.Constant, rcd record.Record) {
	panic("implement me")
}

func (BTreeIndex) Close() {
	panic("implement me")
}
