package query

type ProductScan struct {
	scan1 Scan
	scan2 Scan
}

func NewProductScan(scan1 Scan, scan2 Scan) *ProductScan {
	return &ProductScan{
		scan1: scan1,
		scan2: scan2,
	}
}

func (p *ProductScan) BeforeFirst() {
	p.scan1.BeforeFirst()
	p.scan1.Next()
	p.scan2.BeforeFirst()
}

func (p *ProductScan) Next() bool {
	if p.scan2.Next() {
		return true
	}
	p.scan2.BeforeFirst()
	return p.scan2.Next() && p.scan1.Next()
}

func (p *ProductScan) GetInt(fieldName string) int {
	if p.scan1.HasField(fieldName) {
		return p.scan1.GetInt(fieldName)
	}
	return p.scan2.GetInt(fieldName)
}

func (p *ProductScan) GetString(fieldName string) string {
	if p.scan1.HasField(fieldName) {
		return p.scan1.GetString(fieldName)
	}
	return p.scan2.GetString(fieldName)
}

func (p *ProductScan) GetVal(fieldName string) Constant {
	if p.scan1.HasField(fieldName) {
		return p.scan1.GetVal(fieldName)
	}
	return p.scan2.GetVal(fieldName)
}

func (p *ProductScan) HasField(fieldName string) bool {
	return p.scan1.HasField(fieldName) || p.scan2.HasField(fieldName)
}

func (p *ProductScan) Close() {
	p.scan1.Close()
	p.scan2.Close()
}
