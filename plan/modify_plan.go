package plan

import "goodb/query"

type ModifyPlan struct {
	plan UpdatePlan
	predicate *query.Predicate
}

func NewModifyPlanFromPlan(plan UpdatePlan) *ModifyPlan {
	return &ModifyPlan{
		plan: plan,
	}
}

func NewModifyPlan(plan UpdatePlan, pred *query.Predicate) *ModifyPlan {
	return &ModifyPlan{
		plan: plan,
		predicate: pred,
	}
}

func (sp *ModifyPlan) Open() *query.ModifyScan {
	scan := sp.plan.Open()
	return query.NewModifyScan(scan, sp.predicate)
}
