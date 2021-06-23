package test

import (
	"goodb/file"
	"goodb/metadata"
	"goodb/record"
	"goodb/server"
	"os"
	"testing"
)

func TestTableManager(t *testing.T) {
	os.RemoveAll(file.DB_DIR_PREFIX)

	db := server.NewGooDb("test")

	tx := db.NewTx()
	tm := metadata.NewTableManager(true, tx)

	schema := record.NewSchema()
	schema.AddIntField("i")
	schema.AddStringField("v", 9)
	tm.CreateTable("TestTable", schema, tx)

	layout := tm.GetLayout("TestTable", tx)

	if len(layout.Schema().Fields()) != 2 {
		t.Errorf("Got %d fields, expect 2", len(layout.Schema().Fields()))
	}

	for _, fldName := range layout.Schema().Fields() {
		if fldName == "i" && layout.Schema().Type(fldName) != record.INTEGER {
			t.Errorf("Got %d, expect %d", layout.Schema().Type(fldName), record.INTEGER)
		}

		if fldName == "v" && layout.Schema().Type(fldName) != record.VARCHAR {
			t.Errorf("Got %d, expect %d", layout.Schema().Type(fldName), record.VARCHAR)
		}
	}

	if layout.SlotSize() != 21 {
		t.Errorf("Got %d, expect 21", layout.SlotSize())
	}
}
