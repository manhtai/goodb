package opt

import (
	"goodb/metadata"
	"goodb/plan"
)

func NewIndexJoinPlan(p plan.Plan, tablePlan plan.UpdatePlan, info *metadata.IndexInfo, field string) plan.Plan {
	panic("Implement me!")
}
