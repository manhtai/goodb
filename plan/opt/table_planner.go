package opt

import (
	"goodb/metadata"
	"goodb/plan"
	"goodb/query"
	"goodb/record"
	"goodb/tx"
)

type TablePlanner struct {
	tablePlan plan.UpdatePlan
	predicate query.Predicate
	schema    record.Schema
	indexes   map[string]*metadata.IndexInfo
	tx        *tx.Transaction
}

func NewTablePlanner(tblName string, pred query.Predicate, tx *tx.Transaction, mgr *metadata.MetadataManager) *TablePlanner {
	tblPlan := plan.NewTablePlan(tx, tblName, mgr)
	return &TablePlanner{
		tablePlan: tblPlan,
		predicate: pred,
		schema:    tblPlan.Schema(),
		indexes:   mgr.GetIndexInfo(tblName, tx),
		tx:        tx,
	}
}

func (tp *TablePlanner) MakeSelectPlan() plan.Plan {
	pln := tp.makeIndexSelect()
	if pln == nil {
		pln = tp.tablePlan
	}
	return tp.addSelectPred(pln)
}

func (tp *TablePlanner) MakeJoinPlan(p plan.Plan) plan.Plan {
	curSchema := p.Schema()
	joinPredicate := tp.predicate.JoinSubPred(tp.schema, curSchema)
	if joinPredicate == nil {
		return nil
	}
	pln := tp.makeIndexJoin(p, curSchema)
	if pln == nil {
		pln = tp.makeProductJoin(p, curSchema)
	}
	return pln
}

func (tp *TablePlanner) MakeProductPlan(p plan.Plan) plan.Plan {
	pln := tp.addSelectPred(tp.tablePlan)
	return plan.NewProductPlan(pln, p)
}

func (tp *TablePlanner) makeIndexSelect() plan.Plan {
	for fldName, _ := range tp.indexes {
		val := tp.predicate.EquatesWithConstant(fldName)
		if val != nil {
			indexInfo := tp.indexes[fldName]
			return NewIndexSelectPlan(tp.tablePlan, indexInfo, val)
		}
	}
	return nil
}

func (tp *TablePlanner) makeIndexJoin(p plan.Plan, s record.Schema) plan.Plan {
	for fldName, _ := range tp.indexes {
		outerField := tp.predicate.EquatesWithField(fldName)
		if outerField != "" && s.HasField(outerField) {
			indexInfo := tp.indexes[fldName]
			pln := NewIndexJoinPlan(p, tp.tablePlan, indexInfo, outerField)
			pln = tp.addSelectPred(pln)
			return tp.addJoinPred(pln, s)
		}
	}
	return nil
}

func (tp *TablePlanner) makeProductJoin(p plan.Plan, s record.Schema) plan.Plan {
	pln := tp.MakeProductPlan(p)
	return tp.addJoinPred(pln, s)
}

func (tp *TablePlanner) addSelectPred(p plan.Plan) plan.Plan {
	pred := tp.predicate.SelectSubPred(tp.schema)
	if pred != nil {
		return plan.NewSelectPlan(p, *pred)
	}
	return p
}

func (tp *TablePlanner) addJoinPred(p plan.Plan, s record.Schema) plan.Plan {
	joinPred := tp.predicate.JoinSubPred(s, tp.schema)
	if joinPred != nil {
		return plan.NewSelectPlan(p, *joinPred)
	}
	return p
}
