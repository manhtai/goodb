package query

type Expression struct {
	value     *Constant
	fieldName string
}

func NewFieldExpression(fieldName string) *Expression {
	return &Expression{
		fieldName: fieldName,
	}
}

func NewConstantExpression(value *Constant) *Expression {
	return &Expression{
		value: value,
	}
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
