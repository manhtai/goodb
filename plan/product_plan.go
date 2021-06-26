package plan

import (
	"goodb/query"
	"goodb/record"
)

type ProductPlan struct {
	plan1  Plan
	plan2  Plan
	schema *record.Schema
}

func NewProductPlan(plan1 Plan, plan2 Plan) *ProductPlan {
	schema := record.NewSchema()
	schema.Add(plan1.Schema())
	schema.Add(plan2.Schema())

	return &ProductPlan{
		plan1:  plan1,
		plan2:  plan2,
		schema: schema,
	}
}

func (pp *ProductPlan) Open() query.Scan {
	return query.NewProductScan(pp.plan1.Open(), pp.plan2.Open())
}

func (pp *ProductPlan) Schema() record.Schema {
	return *pp.schema
}
