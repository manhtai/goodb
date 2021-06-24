package query

type ConstantKind int

const (
	StringConstant ConstantKind = iota
	IntConstant
)

type Constant struct {
	intVal int
	strVal string
	kind   ConstantKind
}

func NewIntConstant(val int) Constant {
	return Constant{intVal: val, kind: IntConstant}
}

func NewStrConstant(val string) Constant {
	return Constant{strVal: val, kind: StringConstant}
}
