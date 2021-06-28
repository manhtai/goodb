package index

import (
	"goodb/query"
	"goodb/record"
)

type Index interface {
	BeforeFirst(key query.Constant)
	Next() bool
	GetRecord() record.Record
	Insert(data query.Constant, rcd record.Record)
	Delete(data query.Constant, rcd record.Record)
	Close()
}
