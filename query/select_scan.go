package query

import "goodb/record"

type SelectScan struct {
	scan Scan
	pred *Predicate
}

func NewSelectScan(scan *Scan, pred *Predicate) *SelectScan {
	return &SelectScan{
		scan: *scan,
		pred: pred,
	}
}

func (s *SelectScan) BeforeFirst() {
	s.scan.BeforeFirst()
}

func (s *SelectScan) Next() bool {
	for ; s.scan.Next(); {
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
	// FIXME: What to do?
}

func (s *SelectScan) SetString(fieldName string, val string) {
	// FIXME: What to do?
}

func (s *SelectScan) Insert() {
	// FIXME: What to do?
}

func (s *SelectScan) Delete() {
	// FIXME: What to do?
}

func (s *SelectScan) GetRecord() *record.Record {
	// FIXME: What to do?
	return &record.Record{}
}

func (s *SelectScan) MoveToRecord(rcd *record.Record) {
	// FIXME: What to do?
}

func (s *SelectScan) SetVal(fieldName string, val *Constant) {
	if val.kind == StringKind {
		s.SetString(fieldName, val.strVal)
	} else {
		s.SetInt(fieldName, val.intVal)
	}
}
