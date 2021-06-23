package query

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
	for s.scan.Next() {
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
