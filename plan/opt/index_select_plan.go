package opt

import (
	"goodb/metadata"
	"goodb/plan"
	"goodb/query"
)

func NewIndexSelectPlan(tablePlan plan.UpdatePlan, info *metadata.IndexInfo, val *query.Constant) plan.Plan {
	panic("Implement me")
}
