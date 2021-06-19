package query

import "goodb/record"

type ModifyScan struct {
	scan UpdateScan
	pred *Predicate
}

func NewModifyScan(scan UpdateScan, pred *Predicate) *ModifyScan {
	return &ModifyScan{
		scan: scan,
		pred: pred,
	}
}

func (s *ModifyScan) BeforeFirst() {
	s.scan.BeforeFirst()
}

func (s *ModifyScan) Next() bool {
	for ; s.scan.Next(); {
		if s.pred.IsSatisfied(s) {
			return true
		}
	}
	return false
}

func (s *ModifyScan) GetInt(fieldName string) int {
	return s.scan.GetInt(fieldName)
}

func (s *ModifyScan) GetString(fieldName string) string {
	return s.scan.GetString(fieldName)
}

func (s *ModifyScan) GetVal(fieldName string) *Constant {
	return s.scan.GetVal(fieldName)
}

func (s *ModifyScan) HasField(fieldName string) bool {
	return s.scan.HasField(fieldName)
}

func (s *ModifyScan) Close() {
	s.scan.Close()
}

func (s *ModifyScan) SetInt(fieldName string, val int) {
	s.scan.SetInt(fieldName, val)
}

func (s *ModifyScan) SetString(fieldName string, val string) {
	s.scan.SetString(fieldName, val)
}

func (s *ModifyScan) Insert() {
	s.scan.Insert()
}

func (s *ModifyScan) Delete() {
	s.scan.Delete()
}

func (s *ModifyScan) GetRecord() *record.Record {
	return s.scan.GetRecord()
}

func (s *ModifyScan) MoveToRecord(rcd *record.Record) {
	s.scan.MoveToRecord(rcd)
}

func (s *ModifyScan) SetVal(fieldName string, val *Constant) {
	if val.kind == StringKind {
		s.SetString(fieldName, val.strVal)
	} else {
		s.SetInt(fieldName, val.intVal)
	}
}