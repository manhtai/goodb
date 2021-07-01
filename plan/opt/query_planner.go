package opt

import (
	"goodb/metadata"
	"goodb/parse"
	"goodb/plan"
	"goodb/tx"
)

type OptQueryPlanner struct {
	tablePlanners []*TablePlanner
	metadataMgr   *metadata.MetadataManager
}

func NewOptQueryPlanner(metadataMgr *metadata.MetadataManager) plan.QueryPlanner {
	return &OptQueryPlanner{
		metadataMgr:   metadataMgr,
		tablePlanners: make([]*TablePlanner, 0),
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
	var bestTbp *TablePlanner
	var bestPlan plan.Plan
	for _, tbp := range planner.tablePlanners {
		pln := tbp.MakeSelectPlan()
		if bestPlan == nil || pln.RecordsOutput() < bestPlan.RecordsOutput() {
			bestTbp = tbp
			bestPlan = pln
		}
	}
	planner.removeTablePlanner(bestTbp)
	return bestPlan
}

func (planner *OptQueryPlanner) getLowestJoinPlan(p plan.Plan) plan.Plan {
	var bestTbp *TablePlanner
	var bestPlan plan.Plan
	for _, tbp := range planner.tablePlanners {
		pln := tbp.MakeJoinPlan(p)
		if pln != nil && (bestPlan == nil || pln.RecordsOutput() < bestPlan.RecordsOutput()) {
			bestTbp = tbp
			bestPlan = pln
		}
	}
	if bestPlan != nil {
		planner.removeTablePlanner(bestTbp)
	}
	return bestPlan
}

func (planner *OptQueryPlanner) getLowestProductPlan(p plan.Plan) plan.Plan {
	var bestTbp *TablePlanner
	var bestPlan plan.Plan
	for _, tbp := range planner.tablePlanners {
		pln := tbp.MakeProductPlan(p)
		if bestPlan == nil || pln.RecordsOutput() < bestPlan.RecordsOutput() {
			bestTbp = tbp
			bestPlan = pln
		}
	}
	planner.removeTablePlanner(bestTbp)
	return bestPlan
}

func (planner *OptQueryPlanner) removeTablePlanner(p *TablePlanner) {
	var index int
	for i, tbp := range planner.tablePlanners {
		if tbp == p {
			index = i
			break
		}
	}
	planner.tablePlanners = append(planner.tablePlanners[:index], planner.tablePlanners[index+1:]...)
}
