package query

type Expression struct {
	val *Constant
	fldName string
}

func (expr *Expression) Eval(scan Scan) *Constant {
	if expr.val != nil {
		return expr.val
	}
	return scan.GetVal(expr.fldName)
}

func (expr *Expression) IsFieldName() bool {
	return expr.fldName != ""
}

func (expr *Expression) AsConstant() *Constant {
	return expr.val
}