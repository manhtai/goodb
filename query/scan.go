package query

import "goodb/record"

type Scan interface {
	BeforeFirst()
	Next() bool
	GetInt(fieldName string) int
	GetString(fieldName string) string
	HasField(fieldName string) bool
	Close()
	GetVal(fieldName string) Constant
}

type UpdateScan interface {
	Scan

	SetInt(fieldName string, val int)
	SetString(fieldName string, val string)
	SetVal(fieldName string, val Constant)
	Insert()
	Delete()
	GetRecord() record.Record
	MoveToRecord(record record.Record)
}
