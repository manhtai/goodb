package btree

import (
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type DirEntry struct {
	val query.Constant
	blockNum int
}

type BTreeDir struct {
	tx *tx.Transaction
	layout record.Layout
	data BTreePage
	filename string
}
