package query

type Predicate struct {
	terms []*Term
}

func (pre *Predicate) IsSatisfied(s Scan) bool {
	for _, term := range pre.terms {
		if !term.IsSatisfied(s) {
			return false
		}
	}
	return true
}
