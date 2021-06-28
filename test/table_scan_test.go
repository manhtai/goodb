package test

import (
	"fmt"
	"goodb/file"
	"goodb/query"
	"goodb/record"
	"goodb/server"
	"os"
	"testing"
)

func TestTableScan(t *testing.T) {
	os.RemoveAll(file.DB_DIR_PREFIX)

	db := server.NewGooDbBasic("test")
	tx := db.NewTx()

	schema := record.NewSchema()
	schema.AddIntField("i")
	schema.AddStringField("v", 9)

	layout := record.NewLayoutFromSchema(*schema)
	for _, fldName := range layout.Schema().Fields() {
		offset := layout.Offset(fldName)
		fmt.Printf("Field %s has offset %d\n", fldName, offset)
	}

	tableScan := query.NewTableScan(tx, "T", layout)
	total := 50
	count := 0

	for i := 0; i < total; i++ {
		tableScan.Insert()
		tableScan.SetInt("i", i)
		tableScan.SetString("v", fmt.Sprintf("rec %d", i))
		fmt.Printf("Insert %d, 'rec %d' into table\n", i, i)
	}

	tableScan.BeforeFirst()
	for tableScan.Next() {
		i := tableScan.GetInt("i")
		v := tableScan.GetString("v")
		fmt.Printf("Get %d, '%s' from table\n", i, v)
		if v != fmt.Sprintf("rec %d", i) {
			t.Errorf("Expect v = 'rec %d', got %s", i, v)
		}
		count += 1
	}

	if count != total {
		t.Errorf("Expect %d records, got %d", total, count)
	}

	tableScan.Close()
	tx.Commit()
}
