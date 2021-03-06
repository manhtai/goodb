package metadata

import (
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type IndexManager struct {
	layout   record.Layout
	tableMgr TableManager
}

func NewIndexManager(isNew bool, tableMgr TableManager, tx *tx.Transaction) IndexManager {
	if isNew {
		schema := record.NewSchema()
		schema.AddStringField("indexName", MAX_NAME)
		schema.AddStringField("tableName", MAX_NAME)
		schema.AddStringField("fieldName", MAX_NAME)
		tableMgr.CreateTable("idxCat", *schema, tx)
	}

	layout := tableMgr.GetLayout("idxCat", tx)
	return IndexManager{
		layout:   layout,
		tableMgr: tableMgr,
	}
}

func (indexMgr *IndexManager) createIndex(idxName string, tblName string, fldName string, tx *tx.Transaction) {
	idxCatScan := query.NewTableScan(tx, "idxCat", indexMgr.layout)
	idxCatScan.Insert()
	idxCatScan.SetString("indexName", idxName)
	idxCatScan.SetString("tableName", tblName)
	idxCatScan.SetString("fieldName", fldName)
	idxCatScan.Close()
}

func (indexMgr *IndexManager) getIndexInfo(tblName string, tx *tx.Transaction) map[string]*IndexInfo {
	result := make(map[string]*IndexInfo)
	idxCatScan := query.NewTableScan(tx, "idxCat", indexMgr.layout)

	for idxCatScan.Next() {
		if idxCatScan.GetString("tableName") == tblName {
			idxName := idxCatScan.GetString("indexName")
			fldName := idxCatScan.GetString("fieldName")
			tblLayout := indexMgr.tableMgr.GetLayout(tblName, tx)
			indexInfo := &IndexInfo{
				idxName:   idxName,
				fldName:   fldName,
				tblSchema: *tblLayout.Schema(),
				tx:        tx,
			}
			result[fldName] = indexInfo
		}
	}
	idxCatScan.Close()
	return result
}
