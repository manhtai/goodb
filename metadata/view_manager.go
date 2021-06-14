package metadata

import (
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

const MAX_VIEW_DEF = 100

type ViewManager struct {
	tableMgr *TableManager
}

func NewViewManager(isNew bool, tableMgr *TableManager, tx *tx.Transaction) *ViewManager {
	viewMgr := &ViewManager{tableMgr: tableMgr}
	if isNew {
		schema := &record.Schema{}
		schema.AddStringField("viewName", MAX_NAME)
		schema.AddStringField("viewDef", MAX_VIEW_DEF)
		viewMgr.tableMgr.createTable("viewCat", schema, tx)
	}
	return viewMgr
}

func (viewMgr *ViewManager) createView(viewName string, viewDef string, tx *tx.Transaction) {
	layout := viewMgr.tableMgr.getLayout("viewCat", tx)
	viewCatScan := query.NewTableScan(tx, "viewCat", layout)
	viewCatScan.Insert()
	viewCatScan.SetString("viewName", viewName)
	viewCatScan.SetString("viewDef", viewDef)
	viewCatScan.Close()
}

func (viewMgr *ViewManager) getViewDef(viewName string, tx *tx.Transaction) string {
	var result string
	layout := viewMgr.tableMgr.getLayout("viewCat", tx)
	viewCatScan := query.NewTableScan(tx, "viewCat", layout)
	for ; viewCatScan.Next(); {
		if viewCatScan.GetString("viewName") == viewName {
			result = viewCatScan.GetString("viewDef")
		}
	}
	viewCatScan.Close()
	return result
}
