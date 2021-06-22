package record

type Schema struct {
	fields []string
	info   map[string]*FieldInfo
}

func NewSchema() *Schema {
	return &Schema{
		fields: make([]string, 0),
		info:   make(map[string]*FieldInfo),
	}
}

func (s *Schema) AddSchema(field string, schema Schema) {
	fType := schema.Type(field)
	fLength := schema.Length(field)
	s.AddField(field, fType, fLength)
}

func (s *Schema) AddField(field string, fType int, fLength int) {
	s.fields = append(s.fields, field)
	s.info[field] = &FieldInfo{Type: fType, Length: fLength}
}

func (s *Schema) AddIntField(field string) {
	s.AddField(field, INTEGER, 0)
}

func (s *Schema) AddStringField(field string, length int) {
	s.AddField(field, VARCHAR, length)
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
