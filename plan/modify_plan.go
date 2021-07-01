package plan

import (
	"goodb/query"
	"goodb/record"
)

type ModifyPlan struct {
	plan      UpdatePlan
	predicate query.Predicate
}

func NewModifyPlan(p UpdatePlan, pred query.Predicate) *ModifyPlan {
	return &ModifyPlan{
		plan:      p,
		predicate: pred,
	}
}

func (sp *ModifyPlan) Open() query.Scan {
	return sp.plan.Open()
}

func (sp *ModifyPlan) OpenToUpdate() query.UpdateScan {
	scan := sp.plan.OpenToUpdate()
	return query.NewModifyScan(scan, sp.predicate)
}

func (sp *ModifyPlan) Schema() record.Schema {
	return sp.plan.Schema()
}

func (sp *ModifyPlan) RecordsOutput() int {
	return 0
}
