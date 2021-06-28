package query

import "strings"

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

func (c1 Constant) CompareTo(c2 Constant) int {
	if c1.kind == IntConstant {
		if c1.intVal > c2.intVal {
			return 1
		}
		if c1.intVal < c2.intVal {
			return -1
		}
		return 0
	}
	return strings.Compare(c1.strVal, c2.strVal)
}

func (c Constant) Str() string {
	return c.strVal
}

func (c Constant) Int() int {
	return c.intVal
}
