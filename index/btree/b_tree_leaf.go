package btree

import (
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type BTreeLeaf struct {
	tx *tx.Transaction
	layout record.Layout
	key query.Constant
	data BTreePage
	currentSlot int
	filename string
}
