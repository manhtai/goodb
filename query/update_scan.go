package query

import "goodb/record"

type UpdateScan interface {
	Scan

	SetInt(fieldName string, val int)
	SetString(fieldName string, val string)
	Insert()
	Delete()
	GetRecord() *record.Record
	MoveToRecord(record *record.Record)
}
