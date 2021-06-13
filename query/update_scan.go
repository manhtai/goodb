package query

import "goodb/record"

type UpdateScan interface {
	Scan

	setInt(fieldName string, val int)
	setString(fieldName string, val string)
	insert()
	delete()
	getRecord() *record.Record
	moveToRecord()
}
