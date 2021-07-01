package query

import "goodb/record"

type Predicate struct {
	terms []Term
}

func NewPredicateFromTerms(terms []Term) Predicate {
	return Predicate{terms: terms}
}

func (pre *Predicate) IsSatisfied(s Scan) bool {
	for _, term := range pre.terms {
		if !term.IsSatisfied(s) {
			return false
		}
	}
	return true
}

func (pre *Predicate) EquatesWithConstant(name string) *Constant {
	panic("Implement me")
}

func (pre *Predicate) EquatesWithField(name string) string {
	panic("Implement me")
}

func (pre *Predicate) SelectSubPred(schema record.Schema) *Predicate {
	panic("Implement me")
}

func (pre *Predicate) JoinSubPred(s record.Schema, schema record.Schema) *Predicate {
	panic("Implement me")
}
