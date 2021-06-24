package query

type Expression struct {
	value     Constant
	fieldName string
}

func NewFieldExpression(fieldName string) Expression {
	return Expression{
		fieldName: fieldName,
	}
}

func NewConstantExpression(value Constant) Expression {
	return Expression{
		value: value,
	}
}

func (expr *Expression) Eval(scan Scan) Constant {
	if expr.fieldName != "" {
		return scan.GetVal(expr.fieldName)
	}
	return expr.value
}

func (expr *Expression) IsFieldName() bool {
	return expr.fieldName != ""
}

func (expr *Expression) AsConstant() Constant {
	return expr.value
}
