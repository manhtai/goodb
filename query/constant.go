package query

type ConstantKind int

const (
	StringKind ConstantKind = iota
	IntKind
)

type Constant struct {
	intVal int
	strVal string
	kind   ConstantKind
}

func NewIntConstant(val int) *Constant {
	return &Constant{intVal: val, kind: IntKind}
}

func NewStrConstant(val string) *Constant {
	return &Constant{strVal: val, kind: StringKind}
}
