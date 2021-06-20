package plan

import (
	"goodb/query"
	"goodb/record"
)

type Plan interface {
	Open() query.Scan
	Schema() record.Schema
}

type UpdatePlan interface {
	Plan

	OpenToUpdate() query.UpdateScan
}
