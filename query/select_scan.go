package query

import (
	"goodb/record"
)

type SelectScan struct {
	scan UpdateScan
	pred *Predicate
}

func (s *SelectScan) BeforeFirst() {
	s.scan.BeforeFirst()
}

func (s *SelectScan) Next() bool {
	for ; s.Next() ; {
		if s.pred.IsSatisfied(s) {
			return true
		}
	}
	return false
}

func (s *SelectScan) GetInt(fieldName string) int {
	return s.GetInt(fieldName)
}

func (s *SelectScan) GetString(fieldName string) string {
	return s.GetString(fieldName)
}

func (s *SelectScan) GetVal(fieldName string) *Constant {
	return s.GetVal(fieldName)
}

func (s *SelectScan) HasField(fieldName string) bool {
	return s.HasField(fieldName)
}

func (s *SelectScan) Close() {
	s.Close()
}

func (s *SelectScan) SetInt(fieldName string, val int) {
	s.scan.SetInt(fieldName, val)
}

func (s *SelectScan) SetString(fieldName string, val string) {
	s.scan.SetString(fieldName, val)
}

func (s *SelectScan) Insert() {
	s.scan.Insert()
}

func (s *SelectScan) Delete() {
	s.scan.Delete()
}

func (s *SelectScan) GetRecord() *record.Record {
	return s.scan.GetRecord()
}

func (s *SelectScan) MoveToRecord(rcd *record.Record) {
	s.scan.MoveToRecord(rcd)
}
