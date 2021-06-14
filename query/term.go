package query

type Term struct {
	left *Expression
	right *Expression
}

func NewTerm(left *Expression, right *Expression) *Term {
	return &Term{
		left: left,
		right: right,
	}
}

func (term *Term) IsSatisfied(scan Scan) bool {
	return term.left.Eval(scan) == term.right.Eval(scan)
}