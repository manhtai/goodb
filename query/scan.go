package query

type Scan interface {
	beforeFirst()
	next() bool
	getInt(fieldName string) int
	getString(fieldName string) string
	hasField(fieldName string) bool
	close()
}
