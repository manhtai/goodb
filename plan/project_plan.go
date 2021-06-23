package plan

import (
	"goodb/query"
	"goodb/record"
)

type ProjectPlan struct {
	p      Plan
	schema *record.Schema
}

func NewProjectPlan(p Plan, fields []string) *ProjectPlan {
	schema := record.NewSchema()
	for _, fieldName := range fields {
		schema.AddSchema(fieldName, p.Schema())
	}
	return &ProjectPlan{
		p:      p,
		schema: schema,
	}
}

func (pp *ProjectPlan) Open() query.Scan {
	scan := pp.p.Open()
	return query.NewProjectScan(scan, pp.schema.Fields())
}

func (pp *ProjectPlan) Schema() record.Schema {
	return *pp.schema
}
