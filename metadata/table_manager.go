package metadata

import (
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

const MAX_NAME = 16

type TableManager struct {
	tblCatLayout *record.Layout
	fldCatLayout *record.Layout
}

func NewTableManager(isNew bool, tx *tx.Transaction) *TableManager {
	tblCatSchema := &record.Schema{}
	tblCatSchema.AddStringField("tblName", MAX_NAME)
	tblCatSchema.AddIntField("slotSize")
	tblCatLayout := record.NewLayoutFromSchema(tblCatSchema)

	fldCatSchema := &record.Schema{}
	fldCatSchema.AddStringField("tblName", MAX_NAME)
	fldCatSchema.AddStringField("fldName", MAX_NAME)
	fldCatSchema.AddIntField("type")
	fldCatSchema.AddIntField("length")
	fldCatSchema.AddIntField("offset")
	fldCatLayout := record.NewLayoutFromSchema(fldCatSchema)

	tblMgr := &TableManager{
		tblCatLayout: tblCatLayout,
		fldCatLayout: fldCatLayout,
	}

	if isNew {
		tblMgr.createTable("tblCat", tblCatSchema, tx)
		tblMgr.createTable("fldCat", fldCatSchema, tx)
	}

	return tblMgr
}

func (tableMgr *TableManager) createTable(tblName string, schema *record.Schema, tx *tx.Transaction) {
	layout := record.NewLayoutFromSchema(schema)

	tblCatScan := query.NewTableScan(tx, "tblCat", tableMgr.tblCatLayout)
	tblCatScan.Insert()
	tblCatScan.SetString("tblName", tblName)
	tblCatScan.SetInt("slotSize", layout.SlotSize())
	tblCatScan.Close()

	fldCatScan := query.NewTableScan(tx, "fldCat", tableMgr.tblCatLayout)
	for _, fieldName := range schema.Fields() {
		fldCatScan.Insert()
		fldCatScan.SetString("tblName", tblName)
		fldCatScan.SetString("fldName", fieldName)
		fldCatScan.SetInt("type", schema.Type(fieldName))
		fldCatScan.SetInt("length", schema.Length(fieldName))
		fldCatScan.SetInt("offset", layout.Offset(fieldName))
	}
	fldCatScan.Close()
}

func (tableMgr *TableManager) getLayout(tableName string, tx *tx.Transaction) *record.Layout {
	size := -1

	tblCatScan := query.NewTableScan(tx, "tblCat", tableMgr.tblCatLayout)
	for ; tblCatScan.Next(); {
		if tblCatScan.GetString("tblName") == tableName {
			size = tblCatScan.GetInt("slotSize")
			break
		}
	}
	tblCatScan.Close()

	schema := &record.Schema{}
	offsets := make(map[string]int)
	fldCatScan := query.NewTableScan(tx, "fldCat", tableMgr.fldCatLayout)
	for ; fldCatScan.Next(); {
		if fldCatScan.GetString("tblName") == tableName {
			fldName := fldCatScan.GetString("fldName")
			fldType := fldCatScan.GetInt("type")
			fldLen := fldCatScan.GetInt("length")
			offset := fldCatScan.GetInt("offset")
			offsets[fldName] = offset
			schema.AddField(fldName, fldType, fldLen)
		}
	}
	fldCatScan.Close()

	return record.NewLayout(schema, offsets, size)
}