package query

type Scan interface {
	BeforeFirst()
	Next() bool
	GetInt(fieldName string) int
	GetString(fieldName string) string
	HasField(fieldName string) bool
	Close()
	GetVal(fieldName string) *Constant
}
