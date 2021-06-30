package opt

import (
	"goodb/metadata"
	"goodb/parse"
	"goodb/plan"
	"goodb/tx"
)

type OptQueryPlanner struct {
	tablePlanners []*TablePlanner
	metadataMgr *metadata.MetadataManager
}

func NewOptQueryPlanner(metadataMgr *metadata.MetadataManager) plan.QueryPlanner {
	return &OptQueryPlanner{
		metadataMgr: metadataMgr,
		tablePlanners: make([]*TablePlanner, 1),
	}
}

func (planner *OptQueryPlanner) CreatePlan(stmt parse.SelectStatement, tx *tx.Transaction) plan.Plan {
	for _, tableName := range stmt.Tables {
		tablePlan := NewTablePlanner(tableName, stmt.Predicate, tx, planner.metadataMgr)
		planner.tablePlanners = append(planner.tablePlanners, tablePlan)
	}

	currentPlan := planner.getLowestSelectPlan()
	for len(planner.tablePlanners) > 0 {
		p := planner.getLowestJoinPlan(currentPlan)
		if p != nil {
			currentPlan = p
		} else {
			currentPlan = planner.getLowestProductPlan(currentPlan)
		}
	}

	return plan.NewProjectPlan(currentPlan, stmt.Fields)
}

func (planner *OptQueryPlanner) getLowestSelectPlan() plan.Plan {
	panic("Implement me!")
}

func (planner *OptQueryPlanner) getLowestJoinPlan(p plan.Plan) plan.Plan {
	panic("Implement me!")
}

func (planner *OptQueryPlanner) getLowestProductPlan(p plan.Plan) plan.Plan {
	panic("Implement me!")
}
