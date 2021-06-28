package metadata

import (
	"goodb/index"
	"goodb/index/btree"
	"goodb/record"
	"goodb/tx"
)

type IndexInfo struct {
	idxName   string
	fldName   string
	tx        *tx.Transaction
	tblSchema record.Schema
	idxLayout record.Layout
}

func NewIndexInfo(idxName string, fldName string, tx *tx.Transaction, tblSchema record.Schema) *IndexInfo {
	idxLayout := createIndexLayout(tblSchema, fldName)
	return &IndexInfo{
		idxName:   idxName,
		fldName:   fldName,
		tx:        tx,
		tblSchema: tblSchema,
		idxLayout: idxLayout,
	}
}

func (idxInfo *IndexInfo) Open() index.Index {
	bTreeIndex := btree.NewBTreeIndex(idxInfo.tx, idxInfo.idxName, idxInfo.idxLayout)
	return &bTreeIndex
}

func createIndexLayout(tblSchema record.Schema, fldName string) record.Layout {
	schema := record.NewSchema()
	schema.AddIntField("block")
	schema.AddIntField("id")
	if tblSchema.Type(fldName) == record.INTEGER {
		schema.AddIntField("dataVal")
	} else {
		fldLen := tblSchema.Length(fldName)
		schema.AddStringField("dataVal", fldLen)
	}
	return record.NewLayoutFromSchema(schema)
}
