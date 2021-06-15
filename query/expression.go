package query

type Expression struct {
	value     *Constant
	fieldName string
}

func (expr *Expression) Eval(scan Scan) *Constant {
	if expr.value != nil {
		return expr.value
	}
	return scan.GetVal(expr.fieldName)
}

func (expr *Expression) IsFieldName() bool {
	return expr.fieldName != ""
}

func (expr *Expression) AsConstant() *Constant {
	return expr.value
}
