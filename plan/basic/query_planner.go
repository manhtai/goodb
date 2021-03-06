package basic

import (
	"goodb/metadata"
	"goodb/parse"
	"goodb/plan"
	"goodb/tx"
)

type BasicQueryPlanner struct {
	metadataMgr *metadata.MetadataManager
}

func NewBasicQueryPlanner(metadataMgr *metadata.MetadataManager) plan.QueryPlanner {
	return &BasicQueryPlanner{
		metadataMgr: metadataMgr,
	}
}

func (planner *BasicQueryPlanner) CreatePlan(stmt parse.SelectStatement, tx *tx.Transaction) plan.Plan {
	var plans []plan.Plan
	for _, tableName := range stmt.Tables {
		tablePlan := plan.NewTablePlan(tx, tableName, planner.metadataMgr)
		plans = append(plans, tablePlan)
	}

	if len(plans) == 0 {
		return nil
	}

	p := plans[0]
	for _, nextPlan := range plans[1:] {
		p = plan.NewProductPlan(p, nextPlan)
	}

	p = plan.NewSelectPlan(p, stmt.Predicate)
	p = plan.NewProjectPlan(p, stmt.Fields)

	return p
}
