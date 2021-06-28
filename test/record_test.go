package test

import (
	"fmt"
	"goodb/file"
	"goodb/record"
	"goodb/server"
	"os"
	"testing"
)

func TestRecord(t *testing.T) {
	os.RemoveAll(file.DB_DIR_PREFIX)

	db := server.NewGooDbBasic("test")
	tx := db.NewTx()

	schema := record.NewSchema()
	schema.AddIntField("i")
	schema.AddStringField("v", 9)

	layout := record.NewLayoutFromSchema(*schema)
	block := tx.Append("testFile")
	tx.Pin(block)
	recordPage := record.NewRecordPage(tx, block, layout)
	recordPage.Format()

	slot := recordPage.InsertAfter(-1)
	for slot >= 0 {
		recordPage.SetInt(slot, "i", slot)
		recordPage.SetString(slot, "v", fmt.Sprintf("rec %d", slot))
		fmt.Printf("Insert %d, 'rec %d' into record page\n", slot, slot)

		slot = recordPage.InsertAfter(slot)
	}

	slot = recordPage.NextAfter(-1)
	for slot >= 0 {
		i := recordPage.GetInt(slot, "i")
		v := recordPage.GetString(slot, "v")
		fmt.Printf("Get %d, '%s' from slot %d\n", i, v, slot)

		slot = recordPage.NextAfter(slot)
		if v != fmt.Sprintf("rec %d", i) {
			t.Errorf("Expect 'rec %d', got %s", i, v)
		}
	}

	tx.Unpin(block)
	tx.Commit()
}
