package btree

import (
	"goodb/file"
	"goodb/record"
	"goodb/tx"
)

type BTreePage struct {
	tx *tx.Transaction
	block file.Block
	layout record.Layout
}
