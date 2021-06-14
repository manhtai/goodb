package query

type Constant struct {
	intVal int
	strVal string
}

func NewIntConstant(val int) *Constant {
	return &Constant{intVal: val}
}

func NewStrConstant(val string) *Constant {
	return &Constant{strVal: val}
}
