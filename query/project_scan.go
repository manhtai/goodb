package query

type ProjectScan struct {
	scan   Scan
	fields []string
}

func NewProjectScan(scan Scan, fields []string) *ProjectScan {
	return &ProjectScan{
		scan:   scan,
		fields: fields,
	}
}

func (p *ProjectScan) BeforeFirst() {
	p.scan.BeforeFirst()
}

func (p *ProjectScan) Next() bool {
	return p.scan.Next()
}

func (p *ProjectScan) GetInt(fieldName string) int {
	if p.HasField(fieldName) {
		return p.scan.GetInt(fieldName)
	}
	panic("invalid field name")
}

func (p *ProjectScan) GetString(fieldName string) string {
	if p.HasField(fieldName) {
		return p.scan.GetString(fieldName)
	}
	panic("invalid field name")
}

func (p *ProjectScan) GetVal(fieldName string) *Constant {
	if p.HasField(fieldName) {
		return p.scan.GetVal(fieldName)
	}
	panic("invalid field name")
}

func (p *ProjectScan) HasField(fieldName string) bool {
	for _, field := range p.fields {
		if field == fieldName {
			return true
		}
	}
	return false
}

func (p *ProjectScan) Close() {
	p.scan.Close()
}
