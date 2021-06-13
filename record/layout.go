package record

import "goodb/log"

type Layout struct {
	schema *Schema
	offsets map[string]int
	slotSize int
}

func NewLayoutFromSchema(schema *Schema) *Layout {
	offsets := make(map[string]int)
	pos := log.INT_SIZE
	for _, field := range schema.fields {
		offsets[field] = pos
		pos += schema.Length(field)
	}

	return &Layout{
		schema: schema,
		offsets: offsets,
		slotSize: pos,
	}
}

func (l *Layout) Schema() *Schema {
	return l.schema
}

func (l *Layout) Offset(field string) int {
	return l.offsets[field]
}
