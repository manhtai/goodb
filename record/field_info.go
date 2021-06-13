package record

const (
	INTEGER = iota
	VARCHAR
)

type FieldInfo struct {
	Type   int
	Length int
}
