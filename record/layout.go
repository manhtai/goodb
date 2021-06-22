package record

import "goodb/constant"

type Layout struct {
	schema   *Schema
	offsets  map[string]int
	slotSize int
}

func NewLayout(schema *Schema, offsets map[string]int, slotSize int) *Layout {
	return &Layout{
		schema:   schema,
		offsets:  offsets,
		slotSize: slotSize,
	}
}

func NewLayoutFromSchema(schema *Schema) *Layout {
	offsets := make(map[string]int)
	pos := constant.INT_SIZE
	for _, field := range schema.fields {
		offsets[field] = pos
		pos += schema.LengthInBytes(field)
	}

	return &Layout{
		schema:   schema,
		offsets:  offsets,
		slotSize: pos,
	}
}

func (l *Layout) Schema() *Schema {
	return l.schema
}

func (l *Layout) Offset(field string) int {
	return l.offsets[field]
}

func (l *Layout) SlotSize() int {
	return l.slotSize
}
