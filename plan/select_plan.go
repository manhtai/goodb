package plan

import (
	"goodb/query"
	"goodb/record"
)

type SelectPlan struct {
	plan      Plan
	predicate query.Predicate
}

func NewSelectPlan(p Plan, pred query.Predicate) *SelectPlan {
	return &SelectPlan{
		plan:      p,
		predicate: pred,
	}
}

func (sp *SelectPlan) Open() query.Scan {
	scan := sp.plan.Open()
	return query.NewSelectScan(scan, sp.predicate)
}

func (sp *SelectPlan) Schema() record.Schema {
	return sp.plan.Schema()
}
