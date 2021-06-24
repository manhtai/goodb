package record

import (
	"goodb/constant"
	"goodb/log"
)

type Schema struct {
	fields []string
	info   map[string]FieldInfo
}

func NewSchema() *Schema {
	return &Schema{
		fields: make([]string, 0),
		info:   make(map[string]FieldInfo),
	}
}

func (s *Schema) AddSchema(field string, schema Schema) *Schema {
	fType := schema.Type(field)
	fLength := schema.Length(field)
	s.AddField(field, fType, fLength)
	return s
}

func (s *Schema) AddField(field string, fType int, fLength int) *Schema {
	s.fields = append(s.fields, field)
	s.info[field] = FieldInfo{Type: fType, Length: fLength}
	return s
}

func (s *Schema) AddIntField(field string) *Schema {
	s.AddField(field, INTEGER, 0)
	return s
}

func (s *Schema) AddStringField(field string, length int) *Schema {
	s.AddField(field, VARCHAR, length)
	return s
}

func (s *Schema) HasField(field string) bool {
	_, ok := s.info[field]
	return ok
}

func (s *Schema) Fields() []string {
	return s.fields
}

func (s *Schema) Type(field string) int {
	return s.info[field].Type
}

func (s *Schema) Length(field string) int {
	return s.info[field].Length
}

func (s *Schema) Add(schema Schema) {
	for _, field := range schema.fields {
		s.AddSchema(field, schema)
	}
}

func (s *Schema) LengthInBytes(field string) int {
	switch s.Type(field) {
	case INTEGER:
		return constant.INT_SIZE
	default:
		return log.MaxLength(s.Length(field))
	}
}
