package metadata

import (
	"goodb/record"
	"goodb/tx"
)

type IndexInfo struct {
	idxName   string
	fldName   string
	tx        *tx.Transaction
	tblSchema *record.Schema
	idxLayout *record.Layout
	statInfo  *StatInfo
}
