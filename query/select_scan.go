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
	for ; s.scan.Next() ; {
		if s.pred.IsSatisfied(s) {
			return true
		}
	}
	return false
}

func (s *SelectScan) GetInt(fieldName string) int {
	return s.scan.GetInt(fieldName)
}

func (s *SelectScan) GetString(fieldName string) string {
	return s.scan.GetString(fieldName)
}

func (s *SelectScan) GetVal(fieldName string) *Constant {
	return s.scan.GetVal(fieldName)
}

func (s *SelectScan) HasField(fieldName string) bool {
	return s.scan.HasField(fieldName)
}

func (s *SelectScan) Close() {
	s.scan.Close()
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
